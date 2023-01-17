package capture

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/attempt/application"
	"github.com/scrapnode/scraphook/events"
)

func RegisterCaptureRequest(service *Capture, ctx context.Context) error {
	name := "capture_request"
	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.SCHEDULE_REQUEST}
	queue := fmt.Sprintf("%s_%s", name, service.app.Configs.MsgBus.QueueName)
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, queue, UseCaptureRequest(service.app))
	if err != nil {
		return err
	}

	service.logger.Debugw("registered", "queue_name", queue)
	service.cleanup[name] = cleanup
	return nil
}

func UseCaptureRequest(app *application.App) msgbus.SubscribeFn {
	instrumentName := "capture_request"
	run := application.UseCaptureRequest(app, instrumentName)

	return func(ctx context.Context, event *msgbus.Event) error {
		ctx, span := app.Monitor.Trace(ctx, instrumentName, "msgbus_subscribe")
		ctx = attributes.WithContext(ctx, attributes.Attributes{"event.id": event.Id})
		defer span.End()
		logger := app.Logger.With("event_key", event.Key())

		req := &application.CaptureRequestReq{Event: event}
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
