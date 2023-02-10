package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

func WebhookVerifyExisting(app *App, property string) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			id := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(property).String()
			if id == "" {
				logger.Error("no requested webhook is specified")
				return ctx, errors.New("no requested webhook is specified")
			}

			logger = logger.With("webhook_id", id)
			ok, err := app.Repo.Webhook.VerifyExisting(ws, id)
			if err != nil {
				logger.Errorw("could not verify ownership of the requested webhook.", "error", err.Error())
				return ctx, err
			}

			if !ok {
				msg := fmt.Sprintf("webhook #%s is not exist in your workspace", id)
				logger.Errorw(msg)
				return ctx, errors.New(msg)
			}

			return next(ctx)
		}
	}
}
