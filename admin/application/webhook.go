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

func WebhookVerifyOwnership(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			// if we want to check webhook ownership, we need the webhook id
			id := reflect.ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).
				Elem().FieldByName("Id").String()
			// if the ID is empty, we can get ID from request property WebhookReq
			if id == "" {
				field := reflect.ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).
					Elem().FieldByName("WebhookReq")
				if field.IsValid() {
					id = field.Interface().(WebhookReq).Id
				}
			}

			// if request id is not empty -> update action -> need verifying
			if id != "" {
				ws := ctx.Value(pipeline.CTXKEY_WS).(string)
				ok, err := app.Repo.Webhook.BelongToWorkspace(ws, id)
				if err != nil {
					logger.Errorw("could not check whether webhook is belong to workspace or not", "error", err.Error())
					return ctx, err
				}

				if !ok {
					logger.Error("webhook is not exist in workspace")
					return ctx, errors.New("webhook is not exist in workspace")
				}
			}

			return next(ctx)
		}
	}
}
