package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewEndpointGet(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		// don't need verify whether endpoint is exist or not by EndpointVerifyExisting
		// because if we could not find an endpoint, this pipeline will return not found error by itself
		EndpointGetById(app),
	})
}

type EndpointGetReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

type EndpointGetRes struct {
	Endpoint *entities.Endpoint
}

func EndpointGetById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointGetReq)
			endpoint, err := app.Repo.Endpoint.Get(req.WebhookId, req.Id)
			if err != nil {
				logger.Errorw("could not get endpoint", "error", err.Error())
				return ctx, err
			}

			res := &EndpointGetRes{Endpoint: endpoint}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
