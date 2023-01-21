package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/entities"
)

func UseCaptureResponse(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseCaptureResponseParseEvent(app), app.Monitor, instrumentName, "parse_response"),
		pipeline.UseTracing(UseCaptureResponseMarkRequestAsDone(app), app.Monitor, instrumentName, "mark_request_as_done"),
		pipeline.UseTracing(UseCaptureResponsePut(app), app.Monitor, instrumentName, "put_response"),
	})
}

type CaptureResponseReq struct {
	Event    *msgbus.Event
	Response *entities.Response
}

type CaptureResponseRes struct {
	RequestId string
}

func UseCaptureResponseParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureResponseReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Response); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"request.id": req.Response.Id})
			logger.Debugw("parsed message from event", "message_key", req.Response.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseCaptureResponseMarkRequestAsDone(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureResponseReq)
			logger := app.Logger.With("event_key", req.Event.Key()).With("res_key", req.Response.Key(), "request_id", req.Response.RequestId)

			if err := app.Repo.Request.MarkAsDone(req.Response.RequestId); err != nil {
				logger.Errorw(ErrMarkRequestAsDoneFailed.Error(), "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
func UseCaptureResponsePut(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureResponseReq)
			logger := app.Logger.With("event_key", req.Event.Key()).With("res_key", req.Response.Key())

			if err := app.Repo.Response.Put(req.Response); err != nil {
				logger.Errorw(ErrResponsePutFailed.Error(), "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
