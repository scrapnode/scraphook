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
)

func UseExamineRequest(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(UseExamineRequestParseEvent(app), app.Monitor, instrumentName, "parse_trigger"),
		pipeline.UseTracing(UseExamineRequestScan(app), app.Monitor, instrumentName, "scan_requests"),
		pipeline.UseTracing(UseExamineRequestMarkRequestsAsAttempt(app), app.Monitor, instrumentName, "mark_requests_as_attempt"),
		pipeline.UseTracing(UseExamineRequestFilter(app), app.Monitor, instrumentName, "filter_requests"),
	})
}

type ExamineRequestReq struct {
	Event   *msgbus.Event
	Trigger *entities.RequestTrigger
}

type ExamineRequestRes struct {
	Requests []entities.Request
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

			res := &ExamineRequestRes{Requests: []entities.Request{}}
			query := &repositories.RequestScanQuery{
				ScanQuery: database.ScanQuery{Cursor: "", Limit: app.Configs.Trigger.ScanSize},
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
			res := ctx.Value(pipeline.CTXKEY_RES).(*ExamineRequestRes)
			logger := app.Logger.With("event_key", req.Event.Key())

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
			// @TODO: use redis to filter attempt request if it's exceeded
			return next(ctx)
		}
	}
}
