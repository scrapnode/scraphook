package trigger

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/attempt/application"
)

func UseTriggerRequest(app *application.App) func() {
	run := application.UseTriggerRequest(app)

	return func() {
		ctx := context.Background()
		logger := app.Logger

		req := &application.TriggerRequestReq{
			BucketTemplate: app.Configs.BucketTemplate,
			BucketCount:    app.Configs.Trigger.BucketCount,
		}
		ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Errorw("trigger cronjob got error", "error", err.Error())
			return
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.TriggerRequestRes)
		logger.Debugw("trigger cronjob successfully", "endpoint_count", len(res.Endpoints), "trigger_count", len(res.Triggers))
		return
	}
}
