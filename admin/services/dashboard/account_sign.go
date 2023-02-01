package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AccountServer) Sign(ctx context.Context, req *protos.AccountSignReq) (*protos.AccountSignRes, error) {
	request := &application.AccountSignReq{
		Username: req.Username,
		Password: req.Password,
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.sign(ctx)
	if err != nil {
		server.app.Logger.Errorw("invalid username or password", "error", err.Error())
		return nil, status.Error(codes.InvalidArgument, "invalid username or password")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.AccountSignRes)
	res := &protos.AccountSignRes{
		AccessToken:  response.TokenPair.AccessToken,
		RefreshToken: response.TokenPair.RefreshToken,
	}
	return res, nil
}
