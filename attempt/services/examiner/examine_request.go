package examiner

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/attempt/application"
	"github.com/scrapnode/scraphook/events"
)

func RegisterExamineRequest(service *Examiner, ctx context.Context) error {
	name := "examine_request"
	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.TRIGGER_REQUEST}
	queue := fmt.Sprintf("%s_%s", name, service.app.Configs.MsgBus.QueueName)
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, queue, UseExamineRequest(service.app))
	if err != nil {
		return err
	}

	service.logger.Debugw("registered", "queue_name", queue)
	service.cleanup[name] = cleanup
	return nil
}

func UseExamineRequest(app *application.App) msgbus.SubscribeFn {
	run := application.UseExamineRequest(app)

	return func(ctx context.Context, event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ExamineRequestReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("examiner got error", "error", err.Error())
			return err
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ExamineRequestRes)
		logger.Debugw("examine successfully", "request_count", len(res.Requests))
		return nil
	}
}
