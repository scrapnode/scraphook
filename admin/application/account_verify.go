package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewAccountVerify(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		UseAccountVerifyForRoot(app),
		UseAccountVerifyCheckAccount(app),
	})
}

var TOKEN_TYPE_ROOT = "root"

type AccountVerifyReq struct {
	AccessToken string
	Type        string
}

type AccountVerifyRes struct {
	Account *auth.Account
}

func UseAccountVerifyForRoot(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*AccountVerifyReq)
			logger := app.Logger.With("access_token", req.AccessToken, "token_type", req.Type)
			// not root token, go to next step
			if req.Type != TOKEN_TYPE_ROOT {
				return next(ctx)
			}

			account, err := app.Root.Verify(ctx, req.AccessToken)
			if err != nil {
				logger.Errorw(ErrVerifyFailed.Error(), "error", err.Error())
				return ctx, err
			}

			res := &AccountVerifyRes{Account: account}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

// UseAccountVerifyCheckAccount is where we check whether we could verify access token or not
// because we accept multiple authentication algo (access key, SSO, ...)
// so, we will go through each step and verify request in each pipeline
func UseAccountVerifyCheckAccount(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			res, ok := ctx.Value(pipeline.CTXKEY_RES).(*AccountVerifyRes)
			if !ok {
				return ctx, ErrVerifyFailed
			}

			signed := res != nil && res.Account != nil
			if !signed {
				return ctx, ErrVerifyFailed
			}

			return next(ctx)
		}
	}
}
