package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointRuleServer) List(ctx context.Context, req *protos.EndpointRuleListReq) (*protos.EndpointRuleListRes, error) {
	request := &application.EndpointRuleListReq{
		EndpointId: req.EndpointId,
		Cursor:     req.Cursor,
		Size:       req.Size,
		Search:     req.Search,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.list(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not list endpoint rule", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not list endpoint rule")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointRuleListRes)
	res := &protos.EndpointRuleListRes{
		Data:   []*protos.EndpointRuleRecord{},
		Cursor: response.Cursor,
	}
	for _, edp := range response.EndpointRules {
		// be careful with pointer in loop
		endpoint := edp
		res.Data = append(res.Data, protos.ConvertEndpointRuleToRecord(&endpoint))
	}

	return res, nil
}
