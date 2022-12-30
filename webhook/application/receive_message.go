package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
)

var (
	ErrWebhookNotFound = errors.New("webhook: webhook is not found")
)

func UseReceiveMessage(ctx context.Context, infra *infrastructure.Infra) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		UseReceiveMessageGetWebhook(infra),
		UseReceiveMessagePublishMessage(infra),
	})
}

type ReceiveMessageReq struct {
	WebhookId    string
	WebhookToken string
	Message      *entities.Message
}
type ReceiveMessageRes struct {
	Webhook *entities.Webhook
	PubKey  string
}

func UseReceiveMessageGetWebhook(infra *infrastructure.Infra) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)
			logger := infra.Logger.With("webhook_id", req.WebhookId)

			token, err := infra.Repo.Webhook.GetToken(req.WebhookId, req.WebhookToken)
			if err != nil {
				logger.Errorw(ErrWebhookNotFound.Error(), "error", err.Error())
				return ctx, ErrWebhookNotFound
			}

			res := &ReceiveMessageRes{Webhook: token.Webhook}
			return next(context.WithValue(ctx, pipeline.CTXKEY_RES, res))
		}
	}
}

func UseReceiveMessagePublishMessage(infra *infrastructure.Infra) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*ReceiveMessageRes)

			event := &msgbus.Event{
				Workspace: res.Webhook.WorkspaceId,
				App:       res.Webhook.Id,
				Type:      configs.EVENT_TYPE_MESSAGE,
			}
			// not way to let the error is raised here
			_ = event.SetId()
			// but set data is another story, must handle error
			if err := event.SetData(req.Message); err != nil {
				return ctx, err
			}

			pub, err := infra.MsgBus.Pub(ctx, event)
			if err != nil {
				return ctx, err
			}

			res.PubKey = pub.Key
			return next(context.WithValue(ctx, pipeline.CTXKEY_RES, res))
		}
	}
}
