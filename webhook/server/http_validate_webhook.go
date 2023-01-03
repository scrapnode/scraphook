package server

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/webhook/application"
	"net/http"
)

func UseHttpValidateWebhook(app *application.App) *transport.HttpHandler {
	run := application.UseValidateWebhook(app)

	return &transport.HttpHandler{
		Method: http.MethodGet,
		Path:   "/hooks/:hook_id",
		Handler: func(writer http.ResponseWriter, r *http.Request) {
			params := httprouter.ParamsFromContext(r.Context())
			req := &application.ValidateWebhookReq{
				Id:        params.ByName("hook_id"),
				Token:     r.URL.Query().Get(app.Configs.Validator.VerifyTokenQueryName),
				Challenge: r.URL.Query().Get(app.Configs.Validator.ChallengeQueryName),
			}
			logger := app.Logger.With("method", http.MethodPost, "path", r.RequestURI, "http_name", "validate_webhook")

			ctx := context.WithValue(context.Background(), pipeline.CTXKEY_REQ, req)
			ctx, err := run(ctx)
			if err != nil {
				logger.Error(err)
				if err := transport.WriteErr400(writer, err); err != nil {
					logger.Errorw("could not send json data to client", "error", err.Error())
				}
				return
			}

			res := ctx.Value(pipeline.CTXKEY_RES).(*application.ValidateWebhookRes)
			// if res.challenge is set, sent raw string
			if res.Challenge != "" {
				if err := transport.WriteString(writer, res.Challenge); err != nil {
					logger.Errorw("could not send string data to client", "error", err.Error())
				}
				return
			}

			if err := transport.WriteJSON(writer, res); err != nil {
				logger.Errorw("could not send json data to client", "error", err.Error())
			}
		},
	}
}
