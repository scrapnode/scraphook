package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scrapcore/xsender"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/events"
)

func UseForward(app *App, instrumentName string) pipeline.Pipe {
	send := xsender.New(context.Background(), &xsender.Configs{
		TimeoutInSeconds: 5,
		RetryMax:         3,
	})

	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseForwardParseMessage(app), app.Monitor, instrumentName, "parse_messsage"),
		pipeline.UseTracing(UseForwardSend(app, send), app.Monitor, instrumentName, "send"),
		// optional pipeline
		pipeline.UseTracing(UseForwardNotifyResponse(app), app.Monitor, instrumentName, "notify_response"),
	})
}

type ForwardReq struct {
	Event   *msgbus.Event
	Request *entities.Request
}

type ForwardRes struct {
	Response *entities.Response
}

func UseForwardParseMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ForwardReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Request); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}

			ctx = attributes.WithContext(ctx, attributes.Attributes{"request.id": req.Request.Id})
			logger.Debugw("parsed request from event", "request_key", req.Request.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseForwardSend(app *App, send xsender.Send) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate parsed request
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ForwardReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("request_key", req.Request.Key())

			request := &xsender.Request{
				Uri:    req.Request.Uri,
				Method: req.Request.Method,
				Body:   req.Request.Body,
			}
			if err := request.SetHeaders(req.Request.Headers); err != nil {
				logger.Errorw("could not set request header", "error", err.Error())
				return ctx, err
			}

			response, err := send(request)
			if err != nil {
				logger.Errorw("could not send request", "error", err.Error())
				return ctx, err
			}

			headers, err := response.GetHeaders()
			if err != nil {
				logger.Warnw("parse response headers got error", "error", err.Error())
			}
			res := &ForwardRes{
				Response: &entities.Response{
					WorkspaceId: req.Request.WorkspaceId,
					WebhookId:   req.Request.WebhookId,
					EndpointId:  req.Request.EndpointId,
					RequestId:   req.Request.Id,

					Uri:        response.Uri,
					Status:     response.Status,
					Headers:    headers,
					Body:       response.Body,
					Timestamps: app.Clock.Now().UTC().UnixMilli(),
				},
			}
			res.Response.UseId()
			res.Response.UseTs(app.Configs.BucketTemplate, app.Clock.Now().UTC())

			logger.Debugw("forwared successfully", "status_ok", res.Response.OK(), "status", res.Response.Status)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseForwardNotifyResponse(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ForwardReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*ForwardRes)

			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("response_key", req.Request.Key()).
				With("response_key", res.Response.Key())

			event := &msgbus.Event{
				Workspace: res.Response.WorkspaceId,
				App:       res.Response.WebhookId,
				Type:      events.SCHEDULE_RESPONSE,
				Metadata:  map[string]string{},
			}
			event.UseId()

			if err := event.SetData(res.Response); err != nil {
				logger.Errorw("could not set event data", "error", err.Error())
				// IMPORTANT: we ignore all the error in this pipeline
				// because this flow is optional
				return next(ctx)
			}

			pubres, err := app.MsgBus.Pub(ctx, event)
			if err != nil {
				logger.Errorw("could not publish event", "error", err.Error())
				// IMPORTANT: we ignore all the error in this pipeline
				// because this flow is optional
				return next(ctx)
			}

			logger.Debugw("sent notification", "pubkey", pubres.Key)
			return next(ctx)
		}
	}
}
