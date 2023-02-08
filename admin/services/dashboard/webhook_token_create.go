package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookTokenServer) Create(ctx context.Context, req *protos.WebhookTokenCreateReq) (*protos.WebhookTokenRecord, error) {
	request := &application.WebhookTokenCreateReq{
		WebhookId: req.WebhookId,
		Name:      req.Name,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.create(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not save webhook token", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save webhook token")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookTokenCreateRes)
	res := protos.ConvertWebhookTokenToRecord(response.Token)
	return res, nil
}
