package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/repositories"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookTokenList(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseValidator(),
		pipeline.UseWorkspaceValidator(),
		WebhookVerifyExisting(app, "WebhookId"),
		WebhookTokenListFilter(app),
	})
}

type WebhookTokenListReq struct {
	WebhookId string `validate:"required"`
	Cursor    string
	Size      int32 `validate:"gte=0,lte=100"`
	Search    string
}

type WebhookTokenListRes struct {
	Cursor string
	Tokens []entities.WebhookToken
}

func WebhookTokenListFilter(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookTokenListReq)
			query := &repositories.WebhookTokenListQuery{
				ScanQuery: database.ScanQuery{Cursor: req.Cursor, Size: int(req.Size), Search: req.Search},
				WebhookId: req.WebhookId,
			}
			results, err := app.Repo.WebhookToken.List(query)
			if err != nil {
				logger.Errorw("could not list webhook tokens", "error", err.Error())
				return ctx, err
			}

			res := &WebhookTokenListRes{results.Cursor, results.Data}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
