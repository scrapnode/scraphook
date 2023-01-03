package server

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/application"
	"net/http"
)

func UseHttpReceiveMessage(app *application.App) *transport.HttpHandler {
	run := application.UseReceiveMessage(app)

	return &transport.HttpHandler{
		Method: http.MethodPost,
		Path:   "/hooks/:hook_id",
		Handler: func(writer http.ResponseWriter, r *http.Request) {
			params := httprouter.ParamsFromContext(r.Context())
			req := &application.ReceiveMessageReq{
				Id:    params.ByName("hook_id"),
				Token: r.URL.Query().Get(app.Configs.Validator.VerifyTokenQueryName),
			}
			logger := app.Logger.With("method", http.MethodPost, "path", r.RequestURI, "http_name", "receive_message")

			var body transport.Body
			if err := body.FromHttpRequest(r); err != nil {
				logger.Error(err)
				if err := transport.WriteErr400(writer, err); err != nil {
					logger.Errorw("could not parse body", "error", err.Error())
				}
				return
			}

			var headers transport.Headers
			if err := headers.FromHttpRequest(r); err != nil {
				logger.Error(err)
				if err := transport.WriteErr400(writer, err); err != nil {
					logger.Errorw("could not parse headers", "error", err.Error())
				}
				return
			}
			req.Message = &entities.Message{
				Timestamps: app.Clock.Now().UTC().UnixMilli(),
				Headers:    headers.ToString(),
				Body:       body.ToString(),
				Method:     r.Method,
			}
			req.Message.WithId()
			ctx := context.WithValue(context.Background(), pipeline.CTXKEY_REQ, req)
			ctx, err := run(ctx)
			if err != nil {
				logger.Error(err)
				if err := transport.WriteErr400(writer, err); err != nil {
					logger.Errorw("could not send json data to client", "error", err.Error())
				}
				return
			}

			res := ctx.Value(pipeline.CTXKEY_RES)
			if err := transport.WriteJSON(writer, res); err != nil {
				logger.Errorw("could not send json data to client", "error", err.Error())
			}
		},
	}
}
