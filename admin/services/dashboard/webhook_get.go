package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookServer) Get(ctx context.Context, req *protos.WebhookGetReq) (*protos.WebhookRecord, error) {
	request := &application.WebhookGetReq{
		WebhookReq: application.WebhookReq{Id: req.Id},
		// @TODO: change hardcode
		WithTokenCount: 5,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.get(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not get webhook", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not get webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookGetRes)
	res := protos.ConvertWebhookToRecord(response.Webhook, response.Tokens)
	return res, nil
}
