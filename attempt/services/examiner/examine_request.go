package examiner

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/attempt/application"
	"github.com/scrapnode/scraphook/events"
)

func RegisterExamineRequest(service *Examiner, ctx context.Context) error {
	name := "examine_request"
	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.ATTEMPT_TRIGGER_REQUEST}
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
	instrumentName := "examiner"
	run := application.UseExamineRequest(app, instrumentName)

	return func(ctx context.Context, event *msgbus.Event) error {
		ctx, span := app.Monitor.Trace(ctx, instrumentName, "msgbus_subscribe")
		ctx = attributes.WithContext(ctx, attributes.Attributes{"event.id": event.Id})
		defer span.End()
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ExamineRequestReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			span.KO(err.Error())
			logger.Errorw("examiner got error", "error", err.Error())
			return err
		}

		span.OK("forwarded successfully")
		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ExamineRequestRes)
		logger.Debugw("forwarded successfully", "request_count", len(res.Requests))
		return nil
	}
}
