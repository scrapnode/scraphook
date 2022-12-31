package server

import (
	"context"
	"errors"
	"github.com/scrapnode/scraphook/internal/interfaces"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
)

func New(ctx context.Context, name string) (interfaces.Server, error) {
	cfg := configs.FromContext(ctx)
	infra, err := infrastructure.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if name == "grpc" {
		//@TODO: implement gRPC server
		return nil, errors.New("server: gRPC server is not implemented yet")
	}

	// by default, we will serve HTTP server
	return NewHTTP(ctx, infra), nil
}
