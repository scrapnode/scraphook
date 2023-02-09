package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

type EndpointReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

func EndpointVerifyExisting(app *App, webhookProp, endpointProp string) pipeline.Pipeline {
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
				FieldByName(endpointProp).String()
			if id == "" {
				field := reflect.
					ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
					FieldByName("EndpointReq")
				if field.IsValid() {
					webhookId = field.Interface().(EndpointReq).WebhookId
					id = field.Interface().(EndpointReq).Id
				}
			}

			if id == "" {
				logger.Error("no requested endpoint is specified")
				return ctx, errors.New("no requested endpoint is specified")
			}

			logger = logger.With("webhook_id", webhookId, "endpoint_id", id)
			exist, err := app.Repo.Endpoint.Exist(ws, webhookId, id)
			if err != nil {
				logger.Errorw("could not verify whether endpoint is exist or not", "error", err.Error())
				return ctx, err
			}

			if !exist {
				msg := fmt.Sprintf("endpoint #%s is not exist", id)
				logger.Errorw(msg)
				return ctx, errors.New(msg)
			}

			return next(ctx)
		}
	}
}
