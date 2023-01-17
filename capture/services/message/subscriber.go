package message

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/capture/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	instrumentName := "capture_message"
	run := application.UseCaptureMessage(app, instrumentName)

	return func(ctx context.Context, event *msgbus.Event) error {
		ctx, span := app.Monitor.Trace(ctx, instrumentName, "msgbus_subscribe")
		ctx = attributes.WithContext(ctx, attributes.Attributes{"event.id": event.Id})
		defer span.End()
		logger := app.Logger.With("event_key", event.Key())

		req := &application.CaptureMessageReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			span.KO(err.Error())
			logger.Errorw("capture got error", "error", err.Error())
			return nil
		}
		span.OK("captured successfully")

		logger.Debugw("captured successfully")
		return nil
	}
}
