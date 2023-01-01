package server

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/infrastructure"
	"go.uber.org/zap"
)

type Http struct {
	infra  *infrastructure.Infra
	logger *zap.SugaredLogger

	server transport.Transport
}

func NewHTTP(ctx context.Context, infra *infrastructure.Infra) *Http {
	logger := xlogger.FromContext(ctx).With("pkg", "server.http")
	return &Http{infra: infra, logger: logger}
}

func (server *Http) Start(ctx context.Context) error {
	if err := server.infra.Connect(ctx); err != nil {
		return err
	}

	handlers := []*transport.HttpHandler{
		transport.NewHttpPing(ctx, server.infra.Configs.Configs),
	}
	srv, err := transport.NewHttp(ctx, server.infra.Configs.Http, handlers)
	if err != nil {
		return err
	}

	server.server = srv
	server.logger.Debug("connected")
	return nil
}

func (server *Http) Stop(ctx context.Context) error {
	if server.server != nil {
		if err := server.server.Stop(ctx); err != nil {
			server.logger.Errorw("shutdown http server got error", "error", err.Error())
		}
	}

	if err := server.infra.Disconnect(ctx); err != nil {
		return err
	}

	server.logger.Debug("disconnected")
	return nil
}

func (server *Http) Run(ctx context.Context) error {
	server.logger.Debugw("running", "listen_address", server.infra.Configs.Http.ListenAddress)

	if err := server.server.Run(ctx); err != nil {
		return err
	}
	return nil
}
