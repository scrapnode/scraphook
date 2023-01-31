package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (server *AccountServer) Verify(ctx context.Context, req *protos.VerifyReq) (*protos.VerifyRes, error) {
	request := &application.AccountVerifyReq{
		AccessToken: req.AccessToken,
	}
	if meta, ok := metadata.FromIncomingContext(ctx); ok {
		if types := meta.Get("X-ScrapNode-Token-Type"); len(types) > 0 {
			request.Type = types[0]
		}
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.verify(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not verify account", "error", err.Error())
		return nil, status.Error(codes.InvalidArgument, "could not verify account")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.AccountVerifyRes)
	res := &protos.VerifyRes{
		Workspaces: response.Account.Workspaces,
		Id:         response.Account.Id,
		Name:       response.Account.Name,
		Email:      response.Account.Email,
	}
	return res, nil
}
