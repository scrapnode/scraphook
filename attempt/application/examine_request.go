package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/entities"
)

func UseExamineRequest(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseExamineRequestParseEvent(app), app.Monitor, instrumentName, "parse_trigger"),
	})
}

type ExamineRequestReq struct {
	Event   *msgbus.Event
	Trigger *entities.AttemptTrigger
}

type ExamineRequestRes struct {
	Requests []entities.Request
}

func UseExamineRequestParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Trigger); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"trigger.id": req.Trigger.Id})
			logger.Debugw("parsed trigger from event", "trigger_key", req.Trigger.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}
