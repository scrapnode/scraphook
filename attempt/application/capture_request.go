package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/entities"
)

func UseCaptureRequest(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseCaptureRequestParseEvent(app), app.Monitor, instrumentName, "parse_request"),
		pipeline.UseTracing(UseCaptureRequestPut(app), app.Monitor, instrumentName, "put_request"),
	})
}

type CaptureRequestReq struct {
	Event   *msgbus.Event
	Request *entities.Request
}

type CaptureRequestRes struct {
}

func UseCaptureRequestParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Request); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"request.id": req.Request.Id})
			logger.Debugw("parsed message from event", "message_key", req.Request.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}
func UseCaptureRequestPut(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key()).With("req_key", req.Request.Key())

			if err := app.Repo.Request.Put(req.Request); err != nil {
				logger.Errorw(ErrMessagePutFailed.Error(), "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
