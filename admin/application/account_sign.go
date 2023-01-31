package application

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
)

func NewAccountSign(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseAccountSignAsRoot(app), app.Monitor, instrumentName, "sign_as_root"),
		pipeline.UseTracing(UseAccountSignCheckTokens(app), app.Monitor, instrumentName, "check_tokens"),
	})
}

type AccountSignReq struct {
	Username string
	Password string
}

type AccountSignRes struct {
	AccessToken  string
	RefreshToken string
}

func UseAccountSignAsRoot(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*AccountSignReq)
			logger := app.Logger.With("username", req.Username)
			// not format of access key, go to next step
			if ok := auth.IsAccessKeyPair(req.Username, req.Password); !ok {
				return next(ctx)
			}

			tokens, err := app.Root.Sign(ctx, &auth.SignCreds{Username: req.Username, Password: req.Password})
			if err != nil {
				logger.Errorw(ErrSignFailed.Error(), "error", err.Error())
				return ctx, err
			}

			res := &AccountSignRes{AccessToken: tokens.AccessToken, RefreshToken: tokens.RefreshToken}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

// UseAccountSignCheckTokens is where we check whether we could generate tokens or not
// because we accept multiple authentication algo (access key, SSO, ...)
// so, we will go through each step and sign request in each pipeline
func UseAccountSignCheckTokens(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			res, ok := ctx.Value(pipeline.CTXKEY_RES).(*AccountSignRes)
			if !ok {
				return ctx, ErrSignFailed
			}

			signed := res != nil && res.AccessToken != "" && res.RefreshToken != ""
			if !signed {
				return ctx, ErrSignFailed
			}

			return next(ctx)
		}
	}
}
