package capture

import (
	"context"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/attempt/application"
	"go.uber.org/zap"
)

func New(ctx context.Context, app *application.App) transport.Transport {
	logger := xlogger.FromContext(ctx).With("service", "message")
	return &Capture{app: app, logger: logger}
}

type Capture struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup map[string]func() error
}

func (service *Capture) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.cleanup = map[string]func() error{}
	if err := RegisterCaptureMessage(service, ctx); err != nil {
		return err
	}
	if err := RegisterCaptureRequest(service, ctx); err != nil {
		return err
	}
	if err := RegisterCaptureResponse(service, ctx); err != nil {
		return err
	}

	service.logger.Debug("connected")
	return nil
}

func (service *Capture) Stop(ctx context.Context) error {
	if len(service.cleanup) > 0 {
		for name, cleanup := range service.cleanup {
			if err := cleanup(); err != nil {
				service.logger.Errorw(name+": cleanup subscriber got error", "error", err.Error())
			}
		}
	}

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Capture) Run(ctx context.Context) error {
	service.logger.Debugw("running")
	return nil
}
