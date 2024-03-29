package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookServer) List(ctx context.Context, req *protos.WebhookListReq) (*protos.WebhookListRes, error) {
	request := &application.WebhookListReq{Cursor: req.Cursor, Size: req.Size, Search: req.Search}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.list(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not list webhook", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not list webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookListRes)
	res := &protos.WebhookListRes{
		Data:   []*protos.WebhookRecord{},
		Cursor: response.Cursor,
	}
	for _, wh := range response.Webhooks {
		// be careful with pointer in loop
		webhook := wh
		res.Data = append(res.Data, protos.ConvertWebhookToRecord(&webhook, nil))
	}

	return res, nil
}
