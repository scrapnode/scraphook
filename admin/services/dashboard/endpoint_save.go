package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointServer) Save(ctx context.Context, req *protos.EndpointSaveReq) (*protos.EndpointSaveRes, error) {
	request := &application.EndpointSaveReq{
		WebhookId: req.WebhookId,
		Id:        req.Id,
		Name:      req.Name,
		Uri:       req.Uri,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.save(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not save endpoint", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save endpoint")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointSaveRes)
	res := &protos.EndpointSaveRes{
		Id:        response.Endpoint.Id,
		CreatedAt: response.Endpoint.CreatedAt,
		UpdatedAt: response.Endpoint.UpdatedAt,
	}

	return res, nil
}
