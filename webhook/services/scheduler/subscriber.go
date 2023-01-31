package scheduler

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/events"
	"github.com/scrapnode/scraphook/webhook/application"
)

func RegisterSSubscriber(service *Scheduler, ctx context.Context) error {
	name := "schedule"
	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.MESSAGE}
	queue := fmt.Sprintf("%s_%s", name, service.app.Configs.MsgBus.QueueName)
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, queue, UseSubscriber(service.app))
	if err != nil {
		return err
	}

	service.logger.Debugw("registered", "queue_name", queue)
	service.cleanup[name] = cleanup
	return nil
}

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	run := application.UseScheduleForward(app)

	return func(ctx context.Context, event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ScheduleForwardReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("schedule got error", "error", err.Error())
			return nil
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ScheduleForwardRes)
		var success int
		var fail int
		for _, result := range res.Results {
			if result.Error == "" {
				success++
				continue
			}
			fail++
		}

		logger.Debugw("scheduled successfully", "success_count", success, "fail_count", fail)
		return nil
	}
}
