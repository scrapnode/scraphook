package dashboard

import (
	"context"
	"github.com/scrapnode/scraphook/admin/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AccountServer) Verify(ctx context.Context, req *protos.VerifyReq) (*protos.VerifyRes, error) {
	account, err := server.app.Root.Verify(ctx, req.AccessToken)
	if err != nil {
		server.app.Logger.Errorw("could not verify account", "error", err.Error())
		return nil, status.Error(codes.NotFound, "could not verify account")
	}

	res := &protos.VerifyRes{
		Workspaces: account.Workspaces,
		Id:         account.Id,
		Name:       account.Name,
		Email:      account.Email,
	}
	return res, nil
}
