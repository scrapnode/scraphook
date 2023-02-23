package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *EndpointRuleServer) Save(ctx context.Context, req *protos.EndpointRuleSaveReq) (*protos.EndpointRuleRecord, error) {
	request := &application.EndpointRuleSaveReq{
		EndpointId: req.EndpointId,
		Id:         req.Id,
		Rule:       req.Rule,
		Negative:   req.Negative,
		Priority:   req.Priority,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	var err error
	if req.Id == "" {
		ctx, err = server.create(ctx)
	} else {
		ctx, err = server.update(ctx)
	}
	if err != nil {
		server.app.Logger.Errorw("could not save endpoint rule", "error", err.Error())
		return nil, status.Error(codes.Internal, "could not save endpoint rule")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.EndpointRuleSaveRes)
	res := protos.ConvertEndpointRuleToRecord(response.EndpointRule)
	return res, nil
}
