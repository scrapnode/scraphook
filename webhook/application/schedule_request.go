package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
)

func UseScheduleRequest(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseRecovery(app.Logger),
		UseParseMessage(app),
		UseScheduleRequestGetEndpoints(app),
		UseScheduleRequestBuild(app),
		UseScheduleRequestSend(app),
	})
}

type ValidateScheduleReq struct {
	Event     *msgbus.Event
	Message   *entities.Message
	Endpoints []*entities.Endpoint
}

type ValidateScheduleRes struct {
	PubKey string `json:"pubkey"`
}

func UseParseMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.GetData(&req.Message); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			logger.Debugw("parsed message", "message_id", req.Message.Id)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseScheduleRequestGetEndpoints(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_id", req.Message.Id)

			endpoints, err := app.Repo.Webhook.GetEndpoints(req.Message.WorkspaceId, req.Message.WebhookId)
			if err != nil {
				logger.Errorw(ErrGetEndpointsFail.Error(), "error", err.Error())
				return ctx, err
			}
			if len(endpoints) == 0 {
				logger.Warn(ErrNoEndpoints.Error())
				return ctx, ErrNoEndpoints
			}

			req.Endpoints = endpoints
			logger.Debugw("found endpoints", "endpoint_count", len(req.Endpoints))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseScheduleRequestBuild(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			//logger := app.Logger.
			//	With("event_key", req.Event.Key()).
			//	With("message_id", req.Message.Id)

			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseScheduleRequestSend(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			//logger := app.Logger.With("event_id", req.Event.Id)

			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}
