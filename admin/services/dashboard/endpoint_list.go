package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointServer) List(ctx context.Context, req *protos.EndpointListReq) (*protos.EndpointListRes, error) {
	request := &application.EndpointListReq{
		WebhookId: req.WebhookId,
		Cursor:    req.Cursor,
		Size:      req.Size,
		Search:    req.Search,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.list(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not list endpoint", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not list endpoint")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointListRes)
	res := &protos.EndpointListRes{
		Data:   []*protos.EndpointRecord{},
		Cursor: response.Cursor,
	}
	for _, edp := range response.Endpoints {
		// be careful with pointer in loop
		endpoint := edp
		res.Data = append(res.Data, protos.ConvertEndpointToRecord(&endpoint))
	}

	return res, nil
}
