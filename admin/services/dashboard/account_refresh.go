package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AccountServer) Refresh(ctx context.Context, req *protos.RefreshReq) (*protos.RefreshRes, error) {
	tokens, err := server.app.Root.Refresh(ctx, &auth.Tokens{
		AccessToken:  req.AccessToken,
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		server.app.Logger.Errorw("could not verify account", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not verify account")
	}

	res := &protos.RefreshRes{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	return res, nil
}
