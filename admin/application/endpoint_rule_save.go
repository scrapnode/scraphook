package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewEndpointRuleCreate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		EndpointVerifyExisting(app, "EndpointId"),
		EndpointRuleSavePrepare(app),
		EndpointRuleSavePutToDatabase(app),
	})
}

func NewEndpointRuleUpdate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		EndpointRuleVerifyExisting(app, "Id"),
		EndpointRuleSavePrepare(app),
		EndpointRuleSavePutToDatabase(app),
	})
}

type EndpointRuleSaveReq struct {
	EndpointId string `validate:"required"`
	Id         string
	Rule       string `validate:"required"`
	Negative   bool
	Priority   int32 `validate:"required"`
}

type EndpointRuleSaveRes struct {
	EndpointRule *entities.EndpointRule
}

func EndpointRuleSavePrepare(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointRuleSaveReq)
			rule := &entities.EndpointRule{
				EndpointId: req.EndpointId,
				Id:         req.Id,
				Rule:       req.Rule,
				Negative:   req.Negative,
				Priority:   req.Priority,
				CreatedAt:  app.Clock.Now().UTC().UnixMilli(),
				UpdatedAt:  app.Clock.Now().UTC().UnixMilli(),
			}
			if req.Id == "" {
				rule.UseId()
			}

			res := &EndpointRuleSaveRes{EndpointRule: rule}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func EndpointRuleSavePutToDatabase(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			res := ctx.Value(pipeline.CTXKEY_RES).(*EndpointRuleSaveRes)
			if err := app.Repo.EndpointRule.Save(res.EndpointRule); err != nil {
				logger.Errorw("could not save endpoint rule", "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
