package trigger

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/attempt/application"
)

func UseTriggerRequest(app *application.App) func() {
	instrumentName := "trigger_request"
	run := application.UseTriggerRequest(app, instrumentName)

	return func() {
		ctx, span := app.Monitor.Trace(context.Background(), instrumentName, "cronjob_handler")
		defer span.End()
		logger := app.Logger

		req := &application.TriggerRequestReq{
			BucketTemplate: app.Configs.BucketTemplate,
			BucketCount:    app.Configs.Trigger.BucketCount,
		}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			span.KO(err.Error())
			logger.Errorw("trigger cronjob got error", "error", err.Error())
			return
		}

		span.OK("trigger cronjob successfully")
		res := ctx.Value(pipeline.CTXKEY_RES).(*application.TriggerRequestRes)
		logger.Debugw("trigger cronjob successfully", "endpoint_count", len(res.Endpoints), "trigger_count", len(res.Triggers))
		return
	}
}