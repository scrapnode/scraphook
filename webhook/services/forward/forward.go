package forward

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
)

type Forward struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup map[string]func() error
}

func New(ctx context.Context, app *application.App) *Forward {
	logger := xlogger.FromContext(ctx).With("service", "examiner")
	return &Forward{app: app, logger: logger}
}

func (service *Forward) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.cleanup = map[string]func() error{}
	if err := RegisterSSubscriber(service, ctx); err != nil {
		return err
	}

	service.logger.Debugw("connected")
	return nil
}

func (service *Forward) Stop(ctx context.Context) error {
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

func (service *Forward) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
