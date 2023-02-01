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
		return nil, status.Error(codes.InvalidArgument, "could not list webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookListRes)
	res := &protos.WebhookListRes{
		Data:   []*protos.WebhookRecord{},
		Cursor: response.Cursor,
	}
	for _, webhook := range response.Webhooks {
		res.Data = append(res.Data, &protos.WebhookRecord{
			WorkspaceId: webhook.WorkspaceId,
			Id:          webhook.Id,
			Name:        webhook.Name,
			CreatedAt:   webhook.CreatedAt,
			UpdatedAt:   webhook.UpdatedAt,
		})
	}

	return res, nil
}
