package application

import (
	"context"
	"fmt"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/events"
	"github.com/sourcegraph/conc"
	"regexp"
	"strings"
)

func UseScheduleForward(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseScheduleForwardParseEvent(app), app.Monitor, instrumentName, "parse_message"),
		pipeline.UseTracing(UseScheduleForwardGetEndpoints(app), app.Monitor, instrumentName, "get_endpoints"),
		pipeline.UseTracing(UseScheduleForwardBuildRequests(app), app.Monitor, instrumentName, "build_requests"),
		pipeline.UseTracing(UseScheduleForwardPublishRequests(app), app.Monitor, instrumentName, "publish_requests"),
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

func UseScheduleForwardParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Message); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"message.id": req.Message.Id})
			logger.Debugw("parsed message from event", "message_key", req.Message.Key())
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
			logger.Debugw("found endpoints", "endpoint_count", len(req.Endpoints))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseScheduleForwardBuildRequests(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ScheduleForwardReq)
			logger := app.Logger.
				With("event_key", req.Event.Key()).
				With("message_key", req.Message.Key())

			res := &ScheduleForwardRes{Requests: []*entities.Request{}, Results: []*pipeline.BatchResult{}}
			for _, endpoint := range req.Endpoints {
				if len(endpoint.Rules) == 0 {
					logger.Warnw("endpoint has no rules, ignore all request", "endpoint_id", endpoint.Id)
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
							MessageId:   req.Message.Id,
							Uri:         endpoint.Uri,
							Status:      entities.REQ_STATUS_INIT,
							Method:      req.Message.Method,
							Headers:     req.Message.Headers,
							Body:        req.Message.Body,
						}
						request.UseId()
						request.UseTs(app.Configs.BucketTemplate, app.Clock.Now().UTC())

						res.Requests = append(res.Requests, request)
					}
				}
			}

			// no request to schedule
			if len(res.Requests) == 0 {
				return ctx, nil
			}

			logger.Debugw("schedule requests", "request_count", len(res.Requests))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)

			return next(ctx)
		}
	}
}

func UseScheduleForwardPublishRequests(app *App) pipeline.Pipeline {
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
						Type:      events.SCHEDULE_REQUEST,
						Metadata:  map[string]string{},
					}
					event.UseId()

					if err := event.SetData(request); err != nil {
						logger.Errorw("could not set event data", "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					// let publish an event to let our system knows we have scheduled a forward request
					if _, err := app.MsgBus.Pub(ctx, event); err != nil {
						logger.Errorw("could not publish event", "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					res.Results = append(res.Results, result)
				})
			}
			wg.Wait()

			logger.Debugw("sent requests", "result_count", len(res.Results))
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)

			return next(ctx)
		}
	}
}
