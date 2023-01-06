package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/configs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func UseReceiveMessage(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		pipeline.UseTracing("receive_message", "validator", pipeline.UseValidator()),
		pipeline.UseTracing("receive_message", "get_webhook", UseReceiveMessageGetWebhook(app)),
		pipeline.UseTracing("receive_message", "publish_message", UseReceiveMessagePublishMessage(app)),
	})
}

type ReceiveMessageReq struct {
	Id    string `validate:"required"`
	Token string `validate:"required"`

	Webhook *entities.Webhook
	Message *entities.Message
}
type ReceiveMessageRes struct {
	PubKey string `json:"pubkey"`
}

func UseReceiveMessageGetWebhook(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)
			logger := app.Logger.With("webhook_id", req.Id)

			token, err := app.Repo.Webhook.GetToken(req.Id, req.Token)
			if err != nil {
				logger.Errorw(ErrWebhookNotFound.Error(), "error", err.Error())
				return ctx, ErrWebhookNotFound
			}

			req.Webhook = token.Webhook
			req.Message.WorkspaceId = req.Webhook.WorkspaceId
			req.Message.WebhookId = req.Webhook.Id

			logger.Debugw("webhook.receive_message: found webhook", "workspace_id", req.Message.WorkspaceId)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseReceiveMessagePublishMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ReceiveMessageReq)
			logger := app.Logger.
				With("webhook_id", req.Message.WebhookId).
				With("workspace_id", req.Message.WorkspaceId)

			event := &msgbus.Event{
				Workspace: req.Webhook.WorkspaceId,
				App:       req.Webhook.Id,
				Type:      configs.EVENT_TYPE_MESSAGE,
				Metadata:  map[string]string{},
			}
			// not way to let the error is raised here
			_ = event.SetId()
			// but set data is another story, must handle error
			if err := event.SetData(req.Message); err != nil {
				return ctx, err
			}

			otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(event.Metadata))

			pub, err := app.MsgBus.Pub(ctx, event)
			if err != nil {
				return ctx, err
			}

			res := &ReceiveMessageRes{PubKey: pub.Key}
			logger.Debugw("webhook.receive_message: published event", "pubkey", res.PubKey)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)

			return next(ctx)
		}
	}
}
