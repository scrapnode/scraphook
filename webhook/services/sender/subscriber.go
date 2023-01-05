package sender

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/webhook/application"
)

func UseSubscriber(app *application.App) msgbus.SubscribeFn {
	run := application.UseDoForward(app)

	// @TODO: if the pipeline error is any kind of msgbus err
	// return that error in this subscriber to let msgbus retry it
	return func(event *msgbus.Event) error {
		logger := app.Logger.With("event_key", event.Key())

		req := &application.DoForwardReq{Event: event}
		ctx := context.WithValue(context.Background(), pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("do.forward: schedule got error", "error", err.Error())
			return nil
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.DoForwardRes)
		logger.Debugw("do.forward: forwarded successfully", "response_key", res.Response.Key())
		return nil
	}
}
