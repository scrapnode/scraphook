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
			logger.Errorw("scheduler: schedule request got error", "error", err.Error())
			return nil
		}

		logger.Debugw("scheduler: schedule successfully")
		return nil
	}
}
