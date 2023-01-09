package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xsender"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/events"
)

func UseDoForward(app *App) pipeline.Pipe {
	send := xsender.New(context.Background(), &xsender.Configs{
		TimeoutInSeconds: 5,
		RetryMax:         3,
	})

	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), &pipeline.TracingConfigs{TraceName: "do_forward", SpanName: "init"}),
		pipeline.UseTracing(UseDoForwardParseMessage(app), &pipeline.TracingConfigs{TraceName: "do_forward", SpanName: "parse_message"}),
		pipeline.UseTracing(UseDoForwardSend(app, send), &pipeline.TracingConfigs{TraceName: "do_forward", SpanName: "send"}),
		// optional pipeline
		pipeline.UseTracing(UseDoForwardNotifyResponse(app), &pipeline.TracingConfigs{TraceName: "do_forward", SpanName: "notify_response"}),
	})
}

type DoForwardReq struct {
	Event   *msgbus.Event
	Request *entities.Request
}

type DoForwardRes struct {
	Response *entities.Response
}

func UseDoForwardParseMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*DoForwardReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.GetData(&req.Request); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			logger.Debugw("do.forward: parsed request from event", "request_key", req.Request.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseDoForwardSend(app *App, send xsender.Send) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*DoForwardReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("request_key", req.Request.Key())

			request := &xsender.Request{
				Uri:    req.Request.Uri,
				Method: req.Request.Method,
				Body:   req.Request.Body,
			}
			if err := request.SetHeaders(req.Request.Headers); err != nil {
				logger.Warnw("do.forward: could not set request header", "error", err.Error())
				return ctx, err
			}

			response, err := send(request)
			if err != nil {
				return ctx, err
			}

			headers, err := response.GetHeaders()
			if err != nil {
				logger.Warnw("do.forward: parse endpoint call headers got error", "error", err.Error())
			}
			res := &DoForwardRes{
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
			res.Response.WithId()

			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseDoForwardNotifyResponse(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*DoForwardReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*DoForwardRes)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("response_key", req.Request.Key()).
				With("response_key", res.Response.Key())

			event := &msgbus.Event{
				Workspace: res.Response.WorkspaceId,
				App:       res.Response.WebhookId,
				Type:      events.SCHEDULE_RES,
				Metadata:  map[string]string{},
			}
			// not way to let the error is raised here
			_ = event.SetId()

			// but set data is another story, must handle error
			if err := event.SetData(res.Response); err != nil {
				logger.Error("schedule.forward: could not set event data")
				// IMPORTANT: we ignore all the error in this pipeline
				// because this flow is optional
				return next(ctx)
			}

			if _, err := app.MsgBus.Pub(ctx, event); err != nil {
				logger.Errorw("schedule.forward: could not publish event")
				// IMPORTANT: we ignore all the error in this pipeline
				// because this flow is optional
				return next(ctx)
			}

			logger.Debugw("schedule.forward: sent notification",
				"status_ok", res.Response.OK(), "status", res.Response.Status,
			)
			return next(ctx)
		}
	}
}
