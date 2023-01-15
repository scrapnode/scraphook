package application

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/scrapnode/scraphook/entities"
)

func UseValidateWebhook(app *App, instrumentName string) pipeline.Pipe {
	return pipeline.New([]pipeline.Pipeline{
		pipeline.UseMetrics(app.Monitor, instrumentName, "exec_time"),
		pipeline.UseTracing(pipeline.UseRecovery(app.Logger), app.Monitor, instrumentName, "init"),
		pipeline.UseTracing(pipeline.UseValidator(), app.Monitor, instrumentName, "validate"),
		pipeline.UseTracing(UseValidateWebhookCheckToken(app), app.Monitor, instrumentName, "check_token"),
	})
}

type ValidateWebhookReq struct {
	Id        string `validate:"required"`
	Token     string `validate:"required"`
	Challenge string

	WebhookToken *entities.WebhookToken
}
type ValidateWebhookRes struct {
	Challenge  string `validate:"challenge"`
	Timestamps int64  `validate:"timestamps"`
}

func UseValidateWebhookCheckToken(app *App) pipeline.Pipeline {
	return func(next pipeline.Pipe) pipeline.Pipe {
		return func(ctx context.Context) (context.Context, error) {
			req := ctx.Value(pipeline.CTXKEY_REQ).(*ValidateWebhookReq)
			logger := app.Logger.With("webhook_id", req.Id, "webhook_token", utils.Censor(req.Token, 5))

			token, err := app.Repo.Webhook.GetToken(req.Id, req.Token)
			if err != nil {
				logger.Errorw(ErrWebhookTokenInvalid.Error(), "error", err.Error())
				return ctx, ErrWebhookTokenInvalid
			}
			// censor token before return value
			token.Censor()

			req.WebhookToken = token
			ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, req)

			res := &ValidateWebhookRes{Challenge: req.Challenge, Timestamps: app.Clock.Now().UTC().UnixMilli()}
			logger.Debugw("webhook.validate: validated successfully", "timestamps", res.Timestamps)
			ctx = context.WithValue(ctx, pipeline.CTXKEY_RES, res)
			return next(ctx)
		}
	}
}
