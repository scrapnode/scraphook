package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/entities"
)

func UseCaptureMessage(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		UseCaptureMessageParseEvent(app),
		UseCaptureMessagePut(app),
	})
}

type CaptureMessageReq struct {
	Event   *msgbus.Event
	Message *entities.Message
}

type CaptureMessageRes struct {
}

func UseCaptureMessageParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureMessageReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Message); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"message.id": req.Message.Id})
			logger.Debugw("parsed message from event", "message_key", req.Message.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseCaptureMessagePut(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*CaptureMessageReq)
			logger := app.Logger.With("event_key", req.Event.Key()).With("msg_key", req.Message.Key())

			if err := app.Repo.Message.Put(req.Message); err != nil {
				logger.Errorw(ErrMessagePutFailed.Error(), "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
