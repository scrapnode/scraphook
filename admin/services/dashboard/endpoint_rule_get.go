package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointRuleServer) Get(ctx context.Context, req *protos.EndpointRuleGetReq) (*protos.EndpointRuleRecord, error) {
	request := &application.EndpointRuleGetReq{EndpointId: req.EndpointId, Id: req.Id}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.get(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not get endpoint rule", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not get endpoint rule")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointRuleGetRes)
	res := protos.ConvertEndpointRuleToRecord(response.EndpointRule)
	return res, nil
}
