package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewWebhookTokenDelete(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookTokenVerifyExisting(app, "WebhookId", "Id"),
		WebhookTokenDeleteById(app),
	})
}

type WebhookTokenDeleteReq struct {
	WebhookTokenReq
}

type WebhookTokenDeleteRes struct {
}

func WebhookTokenDeleteById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookTokenDeleteReq)
			if err := app.Repo.WebhookToken.Delete(req.WebhookId, req.Id); err != nil {
				logger.Errorw("could not delete webhook token", "error", err.Error())
				return ctx, err
			}

			res := &WebhookTokenDeleteRes{}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
