package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

func WebhookTokenVerifyExisting(app *App, property string) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			id := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(property).String()
			if id == "" {
				logger.Error("no requested webhook token is specified")
				return ctx, errors.New("no requested webhook token is specified")
			}

			logger = logger.With("token_id", id)
			exist, err := app.Repo.WebhookToken.Exist(ws, id)
			if err != nil {
				logger.Errorw("could not verify whether webhook token is exist or not", "error", err.Error())
				return ctx, err
			}

			if !exist {
				msg := fmt.Sprintf("webhook token #%s is not exist", id)
				logger.Errorw(msg)
				return ctx, errors.New(msg)
			}

			return next(ctx)
		}
	}
}
