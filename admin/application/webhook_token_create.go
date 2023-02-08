package application

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"time"
)

func NewWebhookTokenCreate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyOwnership(app, "WebhookId"),
		WebhookTokenCreateGenerate(app),
	})
}

type WebhookTokenCreateReq struct {
	WebhookId string `validate:"required"`
	Name      string
}

type WebhookTokenCreateRes struct {
	Token *entities.WebhookToken
}

func WebhookTokenCreateGenerate(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookTokenCreateReq)
			token := &entities.WebhookToken{
				Name:      fmt.Sprintf("Generated token at %s", app.Clock.Now().UTC().Format(time.RFC3339)),
				WebhookId: req.WebhookId,
				CreatedAt: app.Clock.Now().UTC().UnixMilli(),
			}
			if req.Name != "" {
				token.Name = req.Name
			}
			token.UseId()
			token.UseToken(64)

			if err := app.Repo.WebhookToken.Create(token); err != nil {
				logger.Errorw("could not generate token", "error", err.Error())
				return ctx, err
			}

			res := &WebhookTokenCreateRes{Token: token}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
