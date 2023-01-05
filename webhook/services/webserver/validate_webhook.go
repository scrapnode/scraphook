package webserver

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/webhook/application"
	"net/http"
)

func UseValidateWebhook(app *application.App) []*transport.HttpHandler {
	return []*transport.HttpHandler{
		{
			Method:  http.MethodGet,
			Path:    "/hooks/:hook_id",
			Handler: UseValidateWebhookHandler(app),
		},
	}
}

func UseValidateWebhookHandler(app *application.App) http.HandlerFunc {
	run := application.UseValidateWebhook(app)

	return func(writer http.ResponseWriter, r *http.Request) {
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
			transport.WriteErr400(writer, err)
			return
		}

		res := ctx.Value(pipeline.CTXKEY_RES).(*application.ValidateWebhookRes)
		// if res.challenge is set, sent raw string
		if res.Challenge != "" {
			transport.WriteString(writer, res.Challenge)
			return
		}
		// otherwise return all responses object
		transport.WriteJSON(writer, res)
	}
}
