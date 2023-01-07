package application

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/sourcegraph/conc"
	"regexp"
	"strings"
)

func UseScheduleForward(app *App) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		// @TODO: replace with github.com/sourcegraph/conc
		pipeline.UseRecovery(app.Logger),
		pipeline.UseTracing(UseScheduleForwardParseMessage(app), &pipeline.TracingConfigs{TraceName: "schedule_forward", SpanName: "parse_message"}),
		pipeline.UseTracing(UseScheduleForwardGetEndpoints(app), &pipeline.TracingConfigs{TraceName: "schedule_forward", SpanName: "get_endpoints"}),
		pipeline.UseTracing(UseScheduleForwardBuild(app), &pipeline.TracingConfigs{TraceName: "schedule_forward", SpanName: "build"}),
		pipeline.UseTracing(UseScheduleForwardSend(app), &pipeline.TracingConfigs{TraceName: "schedule_forward", SpanName: "send"}),
	})
}

type ScheduleForwardReq struct {
	Event     *msgbus.Event
	Message   *entities.Message
	Endpoints []*entities.Endpoint
}

type ScheduleForwardRes struct {
	Requests []*entities.Request
	Results  []*pipeline.BatchResult
}

func UseScheduleForwardParseMessage(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.GetData(&req.Message); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			logger.Debugw("schedule.forward: parsed message from event", "message_key", req.Message.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseScheduleForwardGetEndpoints(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_key", req.Message.Key())

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
			logger.Debugw("schedule.forward: found endpoints", "endpoint_count", len(req.Endpoints))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseScheduleForwardBuild(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_key", req.Message.Key())

			res := &ScheduleForwardRes{Requests: []*entities.Request{}, Results: []*pipeline.BatchResult{}}
			for _, endpoint := range req.Endpoints {
				if len(endpoint.Rules) == 0 {
					logger.Warnw("schedule.forward: endpoint has no rules, ignore all request", "endpoint_id", endpoint.Id)
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

			logger.Debugw("schedule.forward: schedule requests", "request_count", len(res.Requests))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)

			return next(ctx)
		}
	}
}

func UseScheduleForwardSend(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*ScheduleForwardRes)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_key", req.Message.Key())

			var wg conc.WaitGroup
			for _, r := range res.Requests {
				// reflect the value here to make sure we have no issue with concurrency
				request := r
				wg.Go(func() {
					result := &pipeline.BatchResult{Key: request.Key()}
					event := &msgbus.Event{
						Workspace: request.WorkspaceId,
						App:       request.WebhookId,
						Type:      configs.EVENT_TYPE_SCHEDULE_REQ,
						Metadata:  map[string]string{},
					}
					// not way to let the error is raised here
					_ = event.SetId()

					// but set data is another story, must handle error
					if err := event.SetData(request); err != nil {
						logger.Errorw("schedule.forward: could not set event data", "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					// let publish an event to let our system knows we have scheduled a forward request
					if _, err := app.MsgBus.Pub(ctx, event); err != nil {
						logger.Errorw("schedule.forward: could not publish event", "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					res.Results = append(res.Results, result)
				})
			}
			wg.Wait()

			logger.Debugw("schedule.forward: sent requests", "result_count", len(res.Results))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)

			return next(ctx)
		}
	}
}
