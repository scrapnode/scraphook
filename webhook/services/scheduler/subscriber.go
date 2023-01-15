package scheduler

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	instrumentName := "schedule_forward"
	run := application.UseScheduleForward(app, instrumentName)

	return func(ctx context.Context, event *msgbus.Event) error {
		ctx, span := app.Monitor.Trace(ctx, instrumentName, "msgbus_subscribe")
		ctx = attributes.WithContext(ctx, attributes.Attributes{"event.id": event.Id})
		defer span.End()
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ScheduleForwardReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			span.KO(err.Error())
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

		span.OK("scheduled successfully")
		logger.Debugw("scheduled successfully", "success_count", success, "fail_count", fail)
		return nil
	}
}
