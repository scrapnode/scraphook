package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/repositories"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookGet(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookGetById(app),
		WebhookGetGetTokens(app),
	})
}

type WebhookGetReq struct {
	Id             string `validate:"required"`
	WithTokenCount int
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
			if req.WithTokenCount > 0 {
				query := &repositories.WebhookTokenListQuery{
					ScanQuery: database.ScanQuery{Size: req.WithTokenCount},
					WebhookId: res.Webhook.Id,
				}
				results, err := app.Repo.WebhookToken.List(query)
				if err != nil {
					logger.Errorw("could not get tokens of webhook", "error", err.Error())
					return ctx, err
				}
				res.Tokens = results.Data
				ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			}

			return next(ctx)
		}
	}
}
