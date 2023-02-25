package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewEndpointRuleGet(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		// don't need verify whether endpoint is exist or not by EndpointVerifyExisting
		// because if we could not find an endpoint, this pipeline will return not found error by itself
		EndpointRuleGetById(app),
	})
}

type EndpointRuleGetReq struct {
	EndpointId string `validate:"required"`
	Id         string `validate:"required"`
}

type EndpointRuleGetRes struct {
	EndpointRule *entities.EndpointRule
}

func EndpointRuleGetById(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointRuleGetReq)
			rule, err := app.Repo.EndpointRule.Get(req.EndpointId, req.Id)
			if err != nil {
				logger.Errorw("could not get endpoint rule", "error", err.Error())
				return ctx, err
			}

			res := &EndpointRuleGetRes{EndpointRule: rule}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
