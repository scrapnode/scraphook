package scheduler

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	run := application.UseScheduleForward(app)

	// @TODO: if the pipeline error is any kind of msgbus err
	// return that error in this subscriber to let msgbus retry it
	return func(ctx context.Context, event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ScheduleForwardReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("scheduler.forward: schedule got error", "error", err.Error())
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

		logger.Debugw("scheduler.forward: schedule successfully", "success_count", success, "fail_count", fail)
		return nil
	}
}
