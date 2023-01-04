package application

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/configs"
	"regexp"
	"strings"
	"sync"
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
	Requests []*entities.Request
	Results  []*pipeline.BatchResult
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

			logger.Debugw("schedule.request: parsed message", "message_id", req.Message.Id)
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
			logger.Debugw("schedule.request: found endpoints", "endpoint_count", len(req.Endpoints))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			return next(ctx)
		}
	}
}

func UseScheduleRequestBuild(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_id", req.Message.Id)

			res := &ValidateScheduleRes{Requests: []*entities.Request{}, Results: []*pipeline.BatchResult{}}
			for _, endpoint := range req.Endpoints {
				if len(endpoint.Rules) == 0 {
					logger.Warnw("schedule.request: endpoint has no rules, ignore all request", "endpoint_id", endpoint.Id)
					continue
				}

				for _, rule := range endpoint.Rules {
					if strings.HasPrefix(rule.Rule, "regex::") {
						pattern := strings.Replace(rule.Rule, "regex::", "", -1)
						// %#v will include the field names, but not the struct type
						match, err := regexp.MatchString(pattern, fmt.Sprintf("%+v", req.Message))
						if err != nil {
							app.Logger.Errorw("could not match rule", "error", err.Error(),
								"endpoint_id", endpoint.Id, "rule_id", rule.Id, "rule", rule.Rule,
							)
							continue
						}

						if !match {
							app.Logger.Debugw("rule is not matched",
								"endpoint_id", endpoint.Id, "rule_id", rule.Id, "rule", rule.Rule,
							)
							continue
						}

						request := &entities.Request{
							WorkspaceId: req.Message.WorkspaceId,
							WebhookId:   req.Message.WebhookId,
							EndpointId:  endpoint.Id,
							Uri:         endpoint.Uri,
							Status:      entities.REQ_STATUS_INIT,
							Method:      req.Message.Method,
							Headers:     req.Message.Headers,
							Body:        req.Message.Body,
						}
						request.WithId()
						res.Requests = append(res.Requests, request)
					}
				}
			}

			// no request to schedule
			if len(res.Requests) == 0 {
				return ctx, nil
			}

			logger.Debugw("schedule.request: schedule requests", "request_count", len(res.Requests))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseScheduleRequestSend(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateScheduleReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*ValidateScheduleRes)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_id", req.Message.Id)

			var wg sync.WaitGroup
			for _, r := range res.Requests {
				wg.Add(1)
				go func(request *entities.Request) {
					defer wg.Done()
					result := &pipeline.BatchResult{Key: request.Key()}
					event := &msgbus.Event{
						Workspace: request.WorkspaceId,
						App:       request.WebhookId,
						Type:      configs.EVENT_TYPE_SCHEDULE_REQ,
					}
					// not way to let the error is raised here
					_ = event.SetId()
					// but set data is another story, must handle error
					if err := event.SetData(request); err != nil {
						logger.Errorw("schedule.request: could not set event data", "event_key", event.Key(), "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					if _, err := app.MsgBus.Pub(ctx, event); err != nil {
						logger.Errorw("schedule.request: could not publish event", "event_key", event.Key(), "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					res.Results = append(res.Results, result)
				}(r)
			}
			wg.Wait()

			logger.Debugw("schedule.request: sent requests", "result_count", len(res.Results))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
