package application

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"time"
)

func NewWebhookCreate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookSavePrepare(app),
		WebhookSavePutToDatabase(app),
		WebhookSaveGenerateTokens(app),
	})
}

func NewWebhookUpdate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyOwnership(app, "Id"),
		WebhookSavePrepare(app),
		WebhookSavePutToDatabase(app),
		WebhookSaveGenerateTokens(app),
	})
}

type WebhookSaveReq struct {
	Id                string
	Name              string `validate:"required"`
	AutoGenerateToken bool
}

type WebhookSaveRes struct {
	Webhook *entities.Webhook
	Tokens  []entities.WebhookToken
}

func WebhookSavePrepare(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveReq)
			webhook := &entities.Webhook{
				WorkspaceId: ws,
				Id:          req.Id,
				Name:        req.Name,
				CreatedAt:   app.Clock.Now().UTC().UnixMilli(),
				UpdatedAt:   app.Clock.Now().UTC().UnixMilli(),
			}
			if req.Id == "" {
				webhook.UseId()
			}

			res := &WebhookSaveRes{Webhook: webhook, Tokens: []entities.WebhookToken{}}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func WebhookSavePutToDatabase(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			res := ctx.Value(pipeline.CTXKEY_RES).(*WebhookSaveRes)
			if err := app.Repo.Webhook.Save(res.Webhook); err != nil {
				logger.Errorw("could not save webhook", "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}

func WebhookSaveGenerateTokens(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveReq)
			isUpdated := req.Id != ""
			if isUpdated || !req.AutoGenerateToken {
				logger.Warnw("ignore generate token step", "is_update", isUpdated, "autho_generate_token", req.AutoGenerateToken)
				return next(ctx)
			}

			res := ctx.Value(pipeline.CTXKEY_RES).(*WebhookSaveRes)
			token := &entities.WebhookToken{
				Name:      fmt.Sprintf("Generated token at %s", app.Clock.Now().UTC().Format(time.RFC3339)),
				WebhookId: res.Webhook.Id,
				CreatedAt: app.Clock.Now().UTC().UnixMilli(),
			}
			token.UseId()
			token.UseToken(64)
			if err := app.Repo.WebhookToken.Create(token); err != nil {
				logger.Errorw("could not generate token", "error", err.Error())
				// IMPORTANT: ignore error because user can generate token by themselves later
				return next(ctx)
			}

			res.Tokens = append(res.Tokens, *token)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
