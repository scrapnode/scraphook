package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *WebhookTokenServer) Delete(ctx context.Context, req *protos.WebhookTokenDeleteReq) (*protos.WebhookTokenDeleteRes, error) {
	request := &application.WebhookTokenDeleteReq{
		WebhookId: req.WebhookId,
		Id:        req.Id,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.delete(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not delete webhook token", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not delete webhook token")
	}

	res := &protos.WebhookTokenDeleteRes{}
	return res, nil
}
