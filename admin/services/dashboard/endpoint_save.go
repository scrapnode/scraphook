package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointServer) Save(ctx context.Context, req *protos.EndpointSaveReq) (*protos.EndpointRecord, error) {
	request := &application.EndpointSaveReq{
		WebhookId: req.WebhookId,
		Id:        req.Id,
		Name:      req.Name,
		Uri:       req.Uri,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	var err error
	if req.Id == "" {
		ctx, err = server.create(ctx)
	} else {
		ctx, err = server.update(ctx)
	}
	if err != nil {
		server.app.Logger.Errorw("could not save endpoint", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save endpoint")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointSaveRes)
	res := protos.ConvertEndpointToRecord(response.Endpoint)
	return res, nil
}
