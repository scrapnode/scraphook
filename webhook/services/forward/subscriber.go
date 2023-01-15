package forward

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	instrumentName := "forward"
	run := application.UseForward(app, instrumentName)

	return func(ctx context.Context, event *msgbus.Event) error {
		ctx, span := app.Monitor.Trace(ctx, instrumentName, "msgbus_subscribe")
		ctx = attributes.WithContext(ctx, attributes.Attributes{"event.id": event.Id})
		defer span.End()
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ForwardReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			span.KO(err.Error())
			logger.Errorw("forward got error", "error", err.Error())
			return err
		}

		span.OK("forwarded successfully")
		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ForwardRes)
		logger.Debugw("forwarded successfully", "response_key", res.Response.Key())
		return nil
	}
}
