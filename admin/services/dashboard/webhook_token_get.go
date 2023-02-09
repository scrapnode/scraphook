package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookTokenServer) Get(ctx context.Context, req *protos.WebhookTokenGetReq) (*protos.WebhookTokenRecord, error) {
	request := &application.WebhookTokenGetReq{
		WebhookTokenReq: application.WebhookTokenReq{WebhookId: req.WebhookId, Id: req.Id},
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.get(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not get webhook", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not get webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookTokenGetRes)
	res := protos.ConvertWebhookTokenToRecord(response.Token)
	return res, nil
}
