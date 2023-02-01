package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookSave(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseValidator(),
		pipeline.UseWorkspaceValidator(),
		WebhookSaveVerifyOwnership(app),
		WebhookSavePrepare(app),
		WebhookSavePutToDatabase(app),
		WebhookSaveGenerateTokens(app),
	})
}

type WebhookSaveReq struct {
	Id                 string
	Name               string `validate:"required"`
	GenerateTokenCount int    `validate:"gte=0,lte=5"`
}

type WebhookSaveRes struct {
	Webhook *entities.Webhook
	Tokens  []entities.WebhookToken
}

func WebhookSaveVerifyOwnership(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveReq)
			// if request id is not empty -> update action -> need verifying
			if req.Id != "" {
				ws := ctx.Value(pipeline.CTXKEY_WS).(string)
				ok, err := app.Repo.Webhook.BelongToWorkspace(req.Id, ws)
				if err != nil {
					logger.Errorw("could not check whether webhook is belong to workspace or not", "error", err.Error())
					return ctx, err
				}

				if !ok {
					logger.Error("webhook is not exist in workspace")
					return ctx, errors.New("webhook is not exist in workspace")
				}
			}

			return next(ctx)
		}
	}
}

func WebhookSavePrepare(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveReq)
			webhook := &entities.Webhook{
				WorkspaceId: ws,
				Name:        req.Name,
				CreatedAt:   app.Clock.Now().UTC().UnixMilli(),
			}
			if req.Id != "" {
				webhook.Id = req.Id
				webhook.UpdatedAt = app.Clock.Now().UTC().UnixMilli()
			} else {
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
			if req.GenerateTokenCount == 0 {
				return next(ctx)
			}

			res := ctx.Value(pipeline.CTXKEY_RES).(*WebhookSaveRes)
			for i := 0; i < req.GenerateTokenCount; i++ {
				token := entities.WebhookToken{
					WebhookId: res.Webhook.Id,
					CreatedAt: app.Clock.Now().UTC().UnixMilli(),
				}
				token.UseId()
				token.UseToken(64)
				res.Tokens = append(res.Tokens, token)
			}

			if err := app.Repo.WebhookToken.Create(&res.Tokens); err != nil {
				logger.Errorw("could not generate tokens", "error", err.Error())
				// IMPORTANT: ignore error because user can generate token by themselves later
			}

			return next(ctx)
		}
	}
}
