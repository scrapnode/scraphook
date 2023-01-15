package forward

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	run := application.UseForward(app)

	return func(ctx context.Context, event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.ForwardReq{Event: event}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("forward got error", "error", err.Error())
			return err
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ForwardRes)
		logger.Debugw("forwarded successfully", "response_key", res.Response.Key())
		return nil
	}
}
