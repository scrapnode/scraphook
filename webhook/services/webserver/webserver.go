package webserver

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
)

type Webserver struct {
	app    *application.App
	logger *zap.SugaredLogger

	server transport.Transport
}

func New(ctx context.Context, app *application.App) *Webserver {
	logger := xlogger.FromContext(ctx).With("service", "webserver")
	return &Webserver{app: app, logger: logger}
}

func (service *Webserver) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	handlers := []*transport.HttpHandler{
		transport.NewHttpPing(ctx, service.app.Configs.Configs),
		UseReceiveMessage(service.app),
		UseValidateWebhook(service.app),
	}
	srv, err := transport.NewHttp(ctx, service.app.Configs.Http, handlers)
	if err != nil {
		return err
	}

	service.server = srv
	service.logger.Debug("connected")
	return nil
}

func (service *Webserver) Stop(ctx context.Context) error {
	if service.server != nil {
		if err := service.server.Stop(ctx); err != nil {
			service.logger.Errorw("shutdown http server got error", "error", err.Error())
		}
	}

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Webserver) Run(ctx context.Context) error {
	service.logger.Debugw("running", "listen_address", service.app.Configs.Http.ListenAddress)

	if err := service.server.Run(ctx); err != nil {
		return err
	}
	return nil
}
