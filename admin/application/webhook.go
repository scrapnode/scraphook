package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

type WebhookReq struct {
	Id string `validate:"required"`
}

func WebhookVerifyOwnership(app *App, property string) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			var id string
			if property != "" {
				id = reflect.
					ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
					FieldByName(property).String()
			}
			if id == "" {
				field := reflect.
					ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
					FieldByName("WebhookReq")
				if field.IsValid() {
					id = field.Interface().(WebhookReq).Id
				}
			}

			if id == "" {
				logger.Error("no requested webhook is specified")
				return ctx, errors.New("no requested webhook is specified")
			}

			logger = logger.With("webhook_id", id)
			// if request id is not empty -> update action -> need verifying
			ok, err := app.Repo.Webhook.VerifyOwnership(ws, id)
			if err != nil {
				logger.Errorw("could not verify ownership of the requested webhook.", "error", err.Error())
				return ctx, err
			}

			if !ok {
				logger.Errorw("ownership check is failed")
				return ctx, errors.New("you do not have the right to access the requested webhook")
			}

			return next(ctx)
		}
	}
}
