package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookTokenServer) List(ctx context.Context, req *protos.WebhookTokenListReq) (*protos.WebhookTokenListRes, error) {
	request := &application.WebhookTokenListReq{
		WebhookId: req.WebhookId,
		Cursor:    req.Cursor,
		Size:      req.Size,
		Search:    req.Search,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.list(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not list webhook", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not list webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookTokenListRes)
	res := &protos.WebhookTokenListRes{
		Data:   []*protos.WebhookTokenRecord{},
		Cursor: response.Cursor,
	}
	for _, tk := range response.Tokens {
		// be careful with pointer in loop
		token := tk
		res.Data = append(res.Data, protos.ConvertWebhookTokenToRecord(&token))
	}

	return res, nil
}
