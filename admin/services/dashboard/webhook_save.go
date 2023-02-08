package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookServer) Save(ctx context.Context, req *protos.WebhookSaveReq) (*protos.WebhookRecord, error) {
	request := &application.WebhookSaveReq{
		Id:                req.Id,
		Name:              req.Name,
		AutoGenerateToken: req.AutoGenerateToken,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.save(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not save webhook", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookSaveRes)
	res := protos.ConvertWebhookToRecord(response.Webhook, response.Tokens)
	return res, nil
}
