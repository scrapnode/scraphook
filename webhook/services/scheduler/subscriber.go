package scheduler

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	run := application.UseScheduleRequest(app)
	return func(event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ValidateScheduleReq{Event: event}
		ctx := context.WithValue(context.Background(), pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("scheduler.request: schedule got error", "error", err.Error())
			return nil
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ValidateScheduleRes)
		var success int
		var fail int
		for _, result := range res.Results {
			if result.Error == "" {
				success++
				continue
			}
			fail++
		}

		logger.Debugw("scheduler: schedule successfully", "success_count", success, "fail_count", fail)
		return nil
	}
}
