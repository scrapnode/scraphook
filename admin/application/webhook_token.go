package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

type WebhookTokenReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

func WebhookTokenVerifyExisting(app *App, webhookProp, tokenProp string) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			webhookId := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(webhookProp).String()
			id := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(tokenProp).String()
			if id == "" {
				field := reflect.
					ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
					FieldByName("WebhookTokenReq")
				if field.IsValid() {
					webhookId = field.Interface().(WebhookTokenReq).WebhookId
					id = field.Interface().(WebhookTokenReq).Id
				}
			}

			if id == "" {
				logger.Error("no requested webhook token is specified")
				return ctx, errors.New("no requested webhook token is specified")
			}

			logger = logger.With("webhook token_id", id)
			exist, err := app.Repo.WebhookToken.Exist(ws, webhookId, id)
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
