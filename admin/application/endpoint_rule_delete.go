package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewEndpointRuleDelete(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		// @TODO: remove pipeline.UseValidator because we need to validate in transport layers
		pipeline.UseValidator(),
		EndpointRuleVerifyExisting(app, "Id"),
		EndpointRuleDeleteById(app),
	})
}

type EndpointRuleDeleteReq struct {
	EndpointId string `validate:"required"`
	Id         string `validate:"required"`
}

type EndpointRuleDeleteRes struct {
}

func EndpointRuleDeleteById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointRuleDeleteReq)
			if err := app.Repo.EndpointRule.Delete(req.EndpointId, req.Id); err != nil {
				logger.Errorw("could not delete endpoint rule", "error", err.Error())
				return ctx, err
			}

			res := &EndpointRuleDeleteRes{}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
