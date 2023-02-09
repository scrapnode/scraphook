package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/pipeline"
	"reflect"
)

type EndpointRuleReq struct {
	EndpointId string `validate:"required"`
	Id         string `validate:"required"`
}

func EndpointRuleVerifyExisting(app *App, endpointProp, tokenProp string) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			ws := ctx.Value(pipeline.CTXKEY_WS).(string)
			account := ctx.Value(pipeline.CTXKEY_ACC).(*auth.Account)
			logger := app.Logger.With("ws_id", ws, "account_id", account.Id)

			endpointId := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(endpointProp).String()
			id := reflect.
				ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
				FieldByName(tokenProp).String()
			if id == "" {
				field := reflect.
					ValueOf(ctx.Value(pipeline.CTXKEY_REQ)).Elem().
					FieldByName("EndpointReq")
				if field.IsValid() {
					endpointId = field.Interface().(EndpointRuleReq).EndpointId
					id = field.Interface().(EndpointRuleReq).Id
				}
			}

			if id == "" {
				logger.Error("no requested endpoint rule is specified")
				return ctx, errors.New("no requested endpoint rule is specified")
			}

			logger = logger.With("endpoint_id", endpointId, "endpoint_rule_id", id)
			exist, err := app.Repo.EndpointRule.Exist(ws, endpointId, id)
			if err != nil {
				logger.Errorw("could not verify whether endpoint rule is exist or not", "error", err.Error())
				return ctx, err
			}

			if !exist {
				msg := fmt.Sprintf("endpoint rule #%s is not exist", id)
				logger.Errorw(msg)
				return ctx, errors.New(msg)
			}

			return next(ctx)
		}
	}
}