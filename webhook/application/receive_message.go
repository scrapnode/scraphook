package application

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/configs"
)

var (
	ErrWebhookNotFound = errors.New("webhook: webhook is not found")
)

func UseReceiveMessage(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseValidator(),
		UseReceiveMessageGetWebhook(app),
		UseReceiveMessagePublishMessage(app),
	})
}

type ReceiveMessageReq struct {
	WebhookId    string `validate:"required"`
	WebhookToken string `validate:"required"`
	Webhook      *entities.Webhook
	Message      *entities.Message
}
type ReceiveMessageRes struct {
	PubKey string `json:"pubkey"`
}

func UseReceiveMessageGetWebhook(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)
			logger := app.Logger.With("webhook_id", req.WebhookId)

			token, err := app.Repo.Webhook.GetToken(req.WebhookId, req.WebhookToken)
			if err != nil {
				logger.Errorw(ErrWebhookNotFound.Error(), "error", err.Error())
				return ctx, ErrWebhookNotFound
			}

			req.Webhook = token.Webhook
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseReceiveMessagePublishMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)

			event := &msgbus.Event{
				Workspace: req.Webhook.WorkspaceId,
				App:       req.Webhook.Id,
				Type:      configs.EVENT_TYPE_MESSAGE,
			}
			// not way to let the error is raised here
			_ = event.SetId()
			// but set data is another story, must handle error
			if err := event.SetData(req.Message); err != nil {
				return ctx, err
			}

			pub, err := app.MsgBus.Pub(ctx, event)
			if err != nil {
				return ctx, err
			}

			res := &ReceiveMessageRes{PubKey: pub.Key}
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
