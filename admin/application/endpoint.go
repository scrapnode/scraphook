package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

type EndpointReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}

func EndpointVerifyOwnership(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			// if we want to check webhook ownership, we need the webhook id
			webhookId := reflect.ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).
				Elem().FieldByName("WebhookId").String()
			id := reflect.ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).
				Elem().FieldByName("Id").String()
			// if the ID is empty, we can get ID from request property EndpointReq
			if id == "" {
				field := reflect.ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).
					Elem().FieldByName("EndpointReq")
				if field.IsValid() {
					webhookId = field.Interface().(EndpointReq).WebhookId
					id = field.Interface().(EndpointReq).Id
				}
			}

			// if request id is not empty -> update action -> need verifying
			if id != "" {
				workspaceId := ctx.Value(pipeline.CTXKEY_WS).(string)
				ok, err := app.Repo.Endpoint.BelongToWorkspace(workspaceId, webhookId, id)
				if err != nil {
					logger.Errorw("could not check whether endpoint is belong to workspace or not", "error", err.Error())
					return ctx, err
				}

				if !ok {
					logger.Error("endpoint is not exist in workspace")
					return ctx, errors.New("endpoint is not exist in workspace")
				}
			}

			return next(ctx)
		}
	}
}
