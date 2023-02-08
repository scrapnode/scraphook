package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/repositories"
	"github.com/scrapnode/scraphook/entities"
)

func NewEndpointList(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseValidator(),
		pipeline.UseWorkspaceValidator(),
		WebhookVerifyOwnership(app, "WebhookId"),
		EndpointListFilter(app),
	})
}

type EndpointListReq struct {
	WebhookId string `validate:"required"`
	Cursor    string
	Size      int32 `validate:"gte=0,lte=100"`
	Search    string
}

type EndpointListRes struct {
	Cursor    string
	Endpoints []entities.Endpoint
}

func EndpointListFilter(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointListReq)
			query := &repositories.EndpointListQuery{
				ScanQuery: database.ScanQuery{Cursor: req.Cursor, Size: int(req.Size), Search: req.Search},
				WebhookId: req.WebhookId,
			}
			results, err := app.Repo.Endpoint.List(query)
			if err != nil {
				logger.Errorw("could not list endpoint", "error", err.Error())
				return ctx, err
			}

			res := &EndpointListRes{results.Cursor, results.Data}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
