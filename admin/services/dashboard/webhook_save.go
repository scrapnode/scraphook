package dashboard

import (
	"context"
	"github.com/samber/lo"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"github.com/scrapnode/scraphook/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookServer) Save(ctx context.Context, req *protos.WebhookSaveReq) (*protos.WebhookSaveRes, error) {
	request := &application.WebhookSaveReq{
		Id:                 req.Id,
		Name:               req.Name,
		GenerateTokenCount: 1,
	}
	if req.TokenCount > 0 {
		request.GenerateTokenCount = int(req.TokenCount)
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.save(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not save webhook", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookSaveRes)
	res := &protos.WebhookSaveRes{
		Id:        response.Webhook.Id,
		CreatedAt: response.Webhook.CreatedAt,
		UpdatedAt: response.Webhook.UpdatedAt,
		Tokens: lo.Map(response.Tokens, func(item entities.WebhookToken, _ int) string {
			return item.Token
		}),
	}

	return res, nil
}
