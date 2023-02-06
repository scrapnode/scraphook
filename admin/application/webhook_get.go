package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookGet(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyOwnership(app, "Id"),
		WebhookGetById(app),
		WebhookGetGetTokens(app),
	})
}

type WebhookGetReq struct {
	WebhookReq
	WithTokens bool
}

type WebhookGetRes struct {
	Webhook *entities.Webhook
	Tokens  []entities.WebhookToken
}

func WebhookGetById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookGetReq)
			webhook, err := app.Repo.Webhook.Get(ws, req.Id)
			if err != nil {
				logger.Errorw("could not get webhook", "error", err.Error())
				return ctx, err
			}

			res := &WebhookGetRes{Webhook: webhook, Tokens: []entities.WebhookToken{}}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func WebhookGetGetTokens(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookGetReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*WebhookGetRes)
			if req.WithTokens {
				tokens, err := app.Repo.WebhookToken.ListByWebhookId(res.Webhook.Id)
				if err != nil {
					logger.Errorw("could not get tokens of webhook", "error", err.Error())
					return ctx, err
				}
				res.Tokens = tokens
				ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			}

			return next(ctx)
		}
	}
}
