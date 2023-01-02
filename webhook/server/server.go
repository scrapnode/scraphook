package server

import (
	"context"
	"errors"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scraphook/webhook/application"
	"github.com/scrapnode/scraphook/webhook/configs"
)

func New(ctx context.Context, transport string) (transport.Transport, error) {
	cfg := configs.FromContext(ctx)
	app, err := application.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if transport == "grpc" {
		//@TODO: implement gRPC server
		return nil, errors.New("server: gRPC server is not implemented yet")
	}

	// by default, we will serve HTTP server
	return NewHTTP(ctx, app), nil
}
