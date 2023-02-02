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

func (server *WebhookServer) Get(ctx context.Context, req *protos.WebhookGetReq) (*protos.WebhookGetRes, error) {
	request := &application.WebhookGetReq{
		WebhookReq: application.WebhookReq{Id: req.Id},
		WithTokens: true,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.get(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not get webhook", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not get webhook")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.WebhookGetRes)
	res := &protos.WebhookGetRes{
		Webhook: &protos.WebhookRecord{
			WorkspaceId: response.Webhook.WorkspaceId,
			Id:          response.Webhook.Id,
			Name:        response.Webhook.Name,
			CreatedAt:   response.Webhook.CreatedAt,
			UpdatedAt:   response.Webhook.UpdatedAt,
			Tokens: lo.Map(response.Tokens, func(item entities.WebhookToken, _ int) *protos.WebhookTokenRecord {
				return &protos.WebhookTokenRecord{
					WebhookId: item.WebhookId,
					Id:        item.Id,
					Token:     item.Token,
					CreatedAt: item.CreatedAt,
				}
			}),
		},
	}

	return res, nil
}
