package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/repositories"
	"github.com/scrapnode/scraphook/entities"
)

func NewWebhookList(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseValidator(),
		pipeline.UseWorkspaceValidator(),
		WebhookListFilter(app),
	})
}

type WebhookListReq struct {
	Cursor string
	Size   int32 `validate:"gte=0,lte=100"`
	Search string
}

type WebhookListRes struct {
	Cursor   string
	Webhooks []entities.Webhook
}

func WebhookListFilter(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*WebhookListReq)
			query := &repositories.WebhookListQuery{
				ScanQuery:   database.ScanQuery{Cursor: req.Cursor, Size: int(req.Size), Search: req.Search},
				WorkspaceId: ws,
			}
			results, err := app.Repo.Webhook.List(query)
			if err != nil {
				logger.Errorw("could not list webhook", "error", err.Error())
				return ctx, err
			}

			res := &WebhookListRes{results.Cursor, results.Data}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
