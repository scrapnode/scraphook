package application

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/events"
	"github.com/sourcegraph/conc"
	"time"
)

func UseTrigger(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseTriggerConstructBuckets(app), app.Monitor, instrumentName, "construct_buckets"),
		pipeline.UseTracing(UseTriggerScanEndpoints(app), app.Monitor, instrumentName, "scan_endpoints"),
		pipeline.UseTracing(UseTriggerBuildTriggers(app), app.Monitor, instrumentName, "build_triggers"),
		pipeline.UseTracing(UseTriggerPublish(app), app.Monitor, instrumentName, "publish"),
	})
}

type TriggerReq struct {
	BucketTemplate string `json:"bucket_template"`
	BucketCount    int    `json:"bucket_count"`
	Buckets        []entities.AttemptTrigger
}

type TriggerRes struct {
	Endpoints []entities.Endpoint
	Triggers  []entities.AttemptTrigger
	Results   []pipeline.BatchResult
}

func UseTriggerConstructBuckets(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*TriggerReq)

			count := req.BucketCount
			start := app.Clock.Now().UTC()
			// example of boundaries: [100-199, 200-299, 300-399]
			for count > 0 {
				end := start.Add(-time.Duration(app.Configs.Trigger.BucketSizeInMinutes) * time.Minute)
				req.Buckets = append(req.Buckets, entities.AttemptTrigger{
					Start: start.UnixMilli() - 1,
					End:   end.UnixMilli(),
				})

				start = end
				count--
			}

			return next(ctx)
		}
	}
}

func UseTriggerScanEndpoints(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*TriggerReq)
			logger := app.Logger.With("bucket_template", req.BucketTemplate)

			res := &TriggerRes{Endpoints: []entities.Endpoint{}, Triggers: []entities.AttemptTrigger{}}
			var cursor string

			for {
				results, err := app.Repo.Endpoint.Scan(&database.ScanQuery{Cursor: cursor, Limit: app.Configs.Trigger.ScanSize})
				if err != nil {
					logger.Errorw("could not scan endpoints", "error", err.Error())
					return context.WithValue(ctx, pipeline.CTXKEY_RES, res), err
				}
				res.Endpoints = append(res.Endpoints, results.Records...)

				// no more records cause we didn't get next cursor,
				// stop here
				if results.Cursor == "" {
					break
				}

				// continue scan with next cursor
				cursor = results.Cursor
			}

			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}

func UseTriggerBuildTriggers(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*TriggerReq)
			res := ctx.Value(pipeline.CTXKEY_RES).(*TriggerRes)

			for _, endpoint := range res.Endpoints {
				for _, bucket := range req.Buckets {
					res.Triggers = append(res.Triggers, entities.AttemptTrigger{
						Start:       bucket.Start,
						End:         bucket.End,
						WorkspaceId: endpoint.WorkspaceId,
						WebhookId:   endpoint.WebhookId,
						EndpointId:  endpoint.Id,
					})
				}
			}

			return next(ctx)
		}
	}
}

func UseTriggerPublish(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			res := ctx.Value(pipeline.CTXKEY_RES).(*TriggerRes)
			logger := app.Logger

			var wg conc.WaitGroup
			for _, tr := range res.Triggers {
				// reflect the value here to make sure we have no issue with concurrency
				trigger := tr
				wg.Go(func() {
					result := pipeline.BatchResult{Key: trigger.Key()}
					event := &msgbus.Event{
						Workspace: trigger.WorkspaceId,
						App:       trigger.WebhookId,
						Type:      events.ATTEMPT_TRIGGER,
						Metadata:  map[string]string{},
					}
					event.UseId()

					if err := event.SetData(trigger); err != nil {
						logger.Errorw("could not set event data", "trigger_key", trigger.Key())
						result.Error = err.Error()
						res.Results = append(res.Results, result)
						return
					}

					// let publish an event to let our system knows we have scheduled a forward request
					if _, err := app.MsgBus.Pub(ctx, event); err != nil {
						logger.Errorw("could not publish event", "trigger_key", trigger.Key())
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
