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

func (server *AccountServer) Refresh(ctx context.Context, req *protos.AccountRefreshReq) (*protos.AccountRefreshRes, error) {
	request := &application.AccountRefreshReq{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	}
	if meta, ok := metadata.FromIncomingContext(ctx); ok {
		if types := meta.Get("X-ScrapNode-Token-Type"); len(types) > 0 {
			request.Type = types[0]
		}
	}
	ctx = context.WithValue(ctx, pipeline.CTXKEY_REQ, request)

	ctx, err := server.refresh(ctx)
	if err != nil {
		server.app.Logger.Errorw("could not refresh access token", "error", err.Error())
		return nil, status.Error(codes.InvalidArgument, "could not refresh access token")
	}

	response := ctx.Value(pipeline.CTXKEY_RES).(*application.AccountRefreshRes)
	res := &protos.AccountRefreshRes{
		AccessToken:  response.TokenPair.AccessToken,
		RefreshToken: response.TokenPair.RefreshToken,
	}
	return res, nil
}
