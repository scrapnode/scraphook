package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewEndpointDelete(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		// @TODO: remove pipeline.UseValidator because we need to validate in transport layers
		pipeline.UseValidator(),
		EndpointVerifyExisting(app, "Id"),
		EndpointDeleteById(app),
	})
}

type EndpointDeleteReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

type EndpointDeleteRes struct {
}

func EndpointDeleteById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointDeleteReq)
			if err := app.Repo.Endpoint.Delete(req.WebhookId, req.Id); err != nil {
				logger.Errorw("could not get endpoint", "error", err.Error())
				return ctx, err
			}

			res := &EndpointDeleteRes{}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
