package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewWebhookDelete(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyExisting(app, "Id"),
		WebhookDeleteById(app),
	})
}

type WebhookDeleteReq struct {
	Id string `validate:"required"`
}

type WebhookDeleteRes struct {
}

func WebhookDeleteById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookDeleteReq)
			if err := app.Repo.Webhook.Delete(ws, req.Id); err != nil {
				logger.Errorw("could not delete webhook", "error", err.Error())
				return ctx, err
			}

			res := &WebhookDeleteRes{}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
