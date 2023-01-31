package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewAccountRefresh(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		UseAccountRefreshForRoot(app),
		UseAccountRefreshCheckTokens(app),
	})
}

type AccountRefreshReq struct {
	AccessToken  string
	RefreshToken string
	Type         string
}

type AccountRefreshRes struct {
	AccessToken  string
	RefreshToken string
}

func UseAccountRefreshForRoot(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*AccountRefreshReq)
			logger := app.Logger.With("access_token", req.AccessToken, "token_type", req.Type)
			// not root token, go to next step
			if req.Type != TOKEN_TYPE_ROOT {
				return next(ctx)
			}

			tokens, err := app.Root.Refresh(ctx, &auth.Tokens{AccessToken: req.AccessToken, RefreshToken: req.RefreshToken})
			if err != nil {
				logger.Errorw(ErrRefreshFailed.Error(), "error", err.Error())
				return ctx, err
			}

			res := &AccountRefreshRes{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

// UseAccountRefreshCheckTokens is where we check whether we could refresh tokens or not
// because we accept multiple authentication algo (access key, SSO, ...)
// so, we will go through each step and refresh token in each pipeline
func UseAccountRefreshCheckTokens(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			res, ok := ctx.Value(pipeline.CTXKEY_RES).(*AccountRefreshRes)
			if !ok {
				return ctx, ErrRefreshFailed
			}

			signed := res != nil && res.AccessToken != "" && res.RefreshToken != ""
			if !signed {
				return ctx, ErrRefreshFailed
			}

			return next(ctx)
		}
	}
}
