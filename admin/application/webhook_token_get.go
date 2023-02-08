package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookTokenGet(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyOwnership(app, "WebhookId"),
		WebhookTokenGetById(app),
	})
}

type WebhookTokenGetReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

type WebhookTokenGetRes struct {
	Token *entities.WebhookToken
}

func WebhookTokenGetById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookTokenGetReq)
			token, err := app.Repo.WebhookToken.Get(req.WebhookId, req.Id)
			if err != nil {
				logger.Errorw("could not get webhook token", "error", err.Error())
				return ctx, err
			}

			res := &WebhookTokenGetRes{Token: token}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
