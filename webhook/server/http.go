package server

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
)

type Http struct {
	app    *application.App
	logger *zap.SugaredLogger

	server transport.Transport
}

func NewHTTP(ctx context.Context, app *application.App) *Http {
	logger := xlogger.FromContext(ctx).With("pkg", "server.http")
	return &Http{app: app, logger: logger}
}

func (server *Http) Start(ctx context.Context) error {
	if err := server.app.Connect(ctx); err != nil {
		return err
	}

	handlers := []*transport.HttpHandler{
		transport.NewHttpPing(ctx, server.app.Configs.Configs),
		UseHttpReceiveMessage(server.app),
	}
	srv, err := transport.NewHttp(ctx, server.app.Configs.Http, handlers)
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

	if err := server.app.Disconnect(ctx); err != nil {
		return err
	}

	server.logger.Debug("disconnected")
	return nil
}

func (server *Http) Run(ctx context.Context) error {
	server.logger.Debugw("running", "listen_address", server.app.Configs.Http.ListenAddress)

	if err := server.server.Run(ctx); err != nil {
		return err
	}
	return nil
}
