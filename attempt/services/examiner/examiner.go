package examiner

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/attempt/application"
	"go.uber.org/zap"
)

type Examiner struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup map[string]func() error
}

func New(ctx context.Context, app *application.App) *Examiner {
	logger := xlogger.FromContext(ctx).With("service", "examiner")
	return &Examiner{app: app, logger: logger}
}

func (service *Examiner) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.cleanup = map[string]func() error{}
	if err := RegisterExamineRequest(service, ctx); err != nil {
		return err
	}

	service.logger.Debugw("connected")
	return nil
}

func (service *Examiner) Stop(ctx context.Context) error {
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

func (service *Examiner) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
