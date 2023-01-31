package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AccountServer) Sign(ctx context.Context, req *protos.SignReq) (*protos.SignRes, error) {
	tokens, err := server.app.Root.Sign(ctx, &auth.SignCreds{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		server.app.Logger.Errorw("invalid username or password", "error", err.Error())
		return nil, status.Error(codes.InvalidArgument, "invalid username or password")
	}

	res := &protos.SignRes{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
	return res, nil
}
