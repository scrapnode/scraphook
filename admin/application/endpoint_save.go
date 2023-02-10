package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func NewEndpointCreate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		WebhookVerifyExisting(app, "WebhookId"),
		EndpointSavePrepare(app),
		EndpointSavePutToDatabase(app),
	})
}

func NewEndpointUpdate(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseWorkspaceValidator(),
		pipeline.UseValidator(),
		EndpointVerifyExisting(app, "Id"),
		EndpointSavePrepare(app),
		EndpointSavePutToDatabase(app),
	})
}

type EndpointSaveReq struct {
	WebhookId string `validate:"required"`
	Id        string
	Name      string `validate:"required"`
	Uri       string `validate:"required"`
}

type EndpointSaveRes struct {
	Endpoint *entities.Endpoint
}

func EndpointSavePrepare(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*EndpointSaveReq)
			endpoint := &entities.Endpoint{
				WebhookId: req.WebhookId,
				Id:        req.Id,
				Name:      req.Name,
				Uri:       req.Uri,
				CreatedAt: app.Clock.Now().UTC().UnixMilli(),
				UpdatedAt: app.Clock.Now().UTC().UnixMilli(),
			}
			if req.Id == "" {
				endpoint.UseId()
			}

			res := &EndpointSaveRes{Endpoint: endpoint}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func EndpointSavePutToDatabase(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			res := ctx.Value(pipeline.CTXKEY_RES).(*EndpointSaveRes)
			if err := app.Repo.Endpoint.Save(res.Endpoint); err != nil {
				logger.Errorw("could not save endpoint", "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}
