package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
)

var (
	ErrWebhookTokenInvalid = errors.New("webhook: webhook token is not valid")
)

func UseValidateWebhook(ctx context.Context, infra *infrastructure.Infra) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		UseValidateWebhookCheckToken(infra),
	})
}

type ValidateWebhookReq struct {
	Id    string
	Token string
}
type ValidateWebhookRes struct {
	WebhookToken *entities.WebhookToken
}

func UseValidateWebhookCheckToken(infra *infrastructure.Infra) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateWebhookReq)
			logger := infra.Logger.With("webhook_id", req.Id, "webhook_token", utils.Censor(req.Token, 5))

			token, err := infra.Repo.Webhook.GetToken(req.Id, req.Token)
			if err != nil {
				logger.Errorw(ErrWebhookTokenInvalid.Error(), "error", err.Error())
				return ctx, ErrWebhookTokenInvalid
			}
			// censor token before return value
			token.Censor()

			res := &ValidateWebhookRes{WebhookToken: token}
			return next(context.WithValue(ctx, pipeline.CTXKEY_RES, res))
		}
	}
}
