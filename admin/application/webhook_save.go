package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookSave(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseValidator(),
		pipeline.UseWorkspaceValidator(),
		WebhookSavePrepare(app),
		WebhookSavePutToDatabase(app),
	})
}

type WebhookSaveReq struct {
	Name string `validate:"required"`
}

type WebhookSaveRes struct {
	Webhook *entities.Webhook
}

func WebhookSavePrepare(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveReq)
			webhook := &entities.Webhook{
				WorkspaceId: ctx.Value(pipeline.CTXKEY_WS).(string),
				Name:        req.Name,
				CreatedAt:   app.Clock.Now().UTC().UnixMilli(),
			}
			webhook.UseId()

			res := &WebhookSaveRes{Webhook: webhook}
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

			res := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookSaveRes)
			if err := app.Repo.Webhook.Save(res.Webhook); err != nil {
				logger.Errorw("could not save webhook", "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
