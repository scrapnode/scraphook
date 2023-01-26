package application

import (
	"context"
	"github.com/samber/lo"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/xmonitor/attributes"
	"github.com/scrapnode/scraphook/attempt/repositories"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/events"
	"github.com/sourcegraph/conc"
)

func UseExamineRequest(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseExamineRequestParseEvent(app), app.Monitor, instrumentName, "parse_trigger"),
		pipeline.UseTracing(UseExamineRequestScan(app), app.Monitor, instrumentName, "scan_requests"),
		pipeline.UseTracing(UseExamineRequestMarkRequestsAsAttempt(app), app.Monitor, instrumentName, "mark_requests_as_attempt"),
		pipeline.UseTracing(UseExamineRequestFilter(app), app.Monitor, instrumentName, "filter_requests"),
		pipeline.UseTracing(UseExamineRequestPublishAttemptRequests(app), app.Monitor, instrumentName, "publish_attempt_requests"),
	})
}

type ExamineRequestReq struct {
	Event   *msgbus.Event
	Trigger *entities.RequestTrigger
}

type ExamineRequestRes struct {
	Requests []entities.Request
	Results  []pipeline.BatchResult
}

func UseExamineRequestParseEvent(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			// @TODO: validate event
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			if err := req.Event.ScanData(&req.Trigger); err != nil {
				logger.Errorw(ErrEventDataInvalid.Error(), "error", err.Error())
				return ctx, err
			}
			// @TODO: validate message

			ctx = attributes.WithContext(ctx, attributes.Attributes{"trigger.id": req.Trigger.Id})
			logger.Debugw("parsed trigger from event", "trigger_key", req.Trigger.Key())
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			return next(ctx)
		}
	}
}

func UseExamineRequestScan(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())

			res := &ExamineRequestRes{Requests: []entities.Request{}, Results: []pipeline.BatchResult{}}
			query := &repositories.RequestScanQuery{
				ScanQuery: database.ScanQuery{Cursor: "", Limit: app.Configs.Trigger.ScanSize},
				Filters:   map[string]string{"endpoint_id": req.Trigger.EndpointId},
				Before:    req.Trigger.End,
				After:     req.Trigger.Start,
			}

			for {
				results, err := app.Repo.Request.Scan(query)
				if err != nil {
					logger.Errorw("could not scan request", "error", err.Error())
					return context.WithValue(ctx, pipeline.CTXKEY_RES, res), err
				}
				res.Requests = append(res.Requests, results.Records...)

				// no more records cause we didn't get next cursor,
				// stop here
				if results.Cursor == "" {
					break
				}

				// continue scan with next cursor
				query.Cursor = results.Cursor
			}

			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseExamineRequestMarkRequestsAsAttempt(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())
			res := ctx.Value(pipeline.CTXKEY_RES).(*ExamineRequestRes)

			// no request to process
			if len(res.Requests) == 0 {
				logger.Warn("no request to mark as attempt")
				return next(ctx)
			}

			ids := lo.Map[entities.Request](res.Requests, func(item entities.Request, _ int) string {
				return item.Id
			})
			if err := app.Repo.Request.MarkAsAttempt(ids); err != nil {
				logger.Errorw(ErrMarkRequestAsDoneFailed.Error(), "error", err.Error())
				return ctx, err
			}

			return next(ctx)
		}
	}
}

func UseExamineRequestFilter(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())
			res := ctx.Value(pipeline.CTXKEY_RES).(*ExamineRequestRes)

			// no request to process
			if len(res.Requests) == 0 {
				logger.Warn("no request to mark as attempt")
				return next(ctx)
			}

			dict := map[string]entities.Request{}
			var results []string
			var wg conc.WaitGroup
			for _, r := range res.Requests {
				// reflect the value here to make sure we have no issue with concurrency
				dict[r.Key()] = r
				request := r
				wg.Go(func() {
					count, err := app.Cache.Incr(ctx, request.Key())
					if err != nil {
						app.Logger.Errorw("could not increase counter of request",
							"error", err.Error(),
							"request_key", request.Key())
						return
					}

					if count >= app.Configs.Examiner.MaxCount {
						app.Logger.Warnw("reached max count of attempt",
							"error", err.Error(),
							"request_key", request.Key(),
							"count", count,
							"max_count", app.Configs.Examiner.MaxCount,
						)
						return
					}
					results = append(results, request.Key())
				})
			}
			wg.Wait()

			// reset request list
			res.Requests = []entities.Request{}
			// add valid request to list
			for _, requestKey := range results {
				res.Requests = append(res.Requests, dict[requestKey])
			}

			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseExamineRequestPublishAttemptRequests(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ExamineRequestReq)
			logger := app.Logger.With("event_key", req.Event.Key())
			res := ctx.Value(pipeline.CTXKEY_RES).(*ExamineRequestRes)

			if len(res.Requests) == 0 {
				logger.Warn("found no request to attempt")
				return next(ctx)
			}

			// @TODO: count metrics manually

			var wg conc.WaitGroup
			for _, r := range res.Requests {
				// reflect the value here to make sure we have no issue with concurrency
				request := r
				wg.Go(func() {
					result := pipeline.BatchResult{Key: request.Key()}
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

					// let publish an event to let our system knows we have scheduled a examiner request
					if _, err := app.MsgBus.Pub(ctx, event); err != nil {
						logger.Errorw("could not publish event", "request_key", request.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					res.Results = append(res.Results, pipeline.BatchResult{})
				})
			}
			wg.Wait()

			return next(ctx)
		}
	}
}
