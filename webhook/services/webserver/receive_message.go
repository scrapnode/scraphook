package webserver

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/entities"
	"github.com/scrapnode/scraphook/webhook/application"
	"net/http"
)

func UseReceiveMessage(app *application.App) []*transport.HttpHandler {
	return []*transport.HttpHandler{
		{
			Method:  http.MethodPost,
			Path:    "/hooks/:hook_id",
			Handler: UseReceiveMessageHandler(app),
		},
		{
			Method:  http.MethodPut,
			Path:    "/hooks/:hook_id",
			Handler: UseReceiveMessageHandler(app),
		},
	}
}

func UseReceiveMessageHandler(app *application.App) http.HandlerFunc {
	run := application.UseReceiveMessage(app)

	return func(writer http.ResponseWriter, r *http.Request) {
		params := httprouter.ParamsFromContext(r.Context())
		req := &application.ReceiveMessageReq{
			Id:    params.ByName("hook_id"),
			Token: r.URL.Query().Get(app.Configs.Validator.VerifyTokenQueryName),
		}
		logger := app.Logger.With("method", http.MethodPost, "path", r.RequestURI, "http_name", "receive_message")

		var body transport.Body
		if err := body.FromHttpRequest(r); err != nil {
			logger.Error(err)
			transport.WriteErr400(writer, err)
			return
		}

		var headers transport.Headers
		if err := headers.FromHttpRequest(r); err != nil {
			logger.Error(err)
			transport.WriteErr400(writer, err)
			return
		}

		req.Message = &entities.Message{
			Timestamps: app.Clock.Now().UTC().UnixMilli(),
			Headers:    headers.ToString(),
			Body:       body.ToString(),
			Method:     r.Method,
		}
		req.Message.UseId()
		req.Message.UseTs(app.Configs.BucketTemplate, app.Clock.Now().UTC())

		ctx := context.WithValue(context.Background(), pipeline.CTXKEY_REQ, req)
		ctx, err := run(ctx)
		if err != nil {
			logger.Error(err)
			transport.WriteErr400(writer, err)
			return
		}

		res := ctx.Value(pipeline.CTXKEY_RES)
		transport.WriteJSON(writer, res)
	}
}
