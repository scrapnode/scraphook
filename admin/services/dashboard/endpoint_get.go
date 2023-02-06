package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointServer) Get(ctx context.Context, req *protos.EndpointGetReq) (*protos.EndpointGetRes, error) {
	request := &application.EndpointGetReq{
		EndpointReq: application.EndpointReq{
			WebhookId: req.WebhookId,
			Id:        req.Id,
		},
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.get(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not get endpoint", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not get endpoint")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointGetRes)
	res := &protos.EndpointGetRes{
		Endpoint: &protos.EndpointRecord{
			WorkspaceId: response.Endpoint.WorkspaceId,
			Id:          response.Endpoint.Id,
			Name:        response.Endpoint.Name,
			Uri:         response.Endpoint.Uri,
			CreatedAt:   response.Endpoint.CreatedAt,
			UpdatedAt:   response.Endpoint.UpdatedAt,
		},
	}

	return res, nil
}
