package scheduler

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
)

type Scheduler struct {
	app    *application.App
	logger *zap.SugaredLogger
}

func New(ctx context.Context, app *application.App) *Scheduler {
	logger := xlogger.FromContext(ctx).With("service", "scheduler")
	return &Scheduler{app: app, logger: logger}
}

func (service *Scheduler) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.logger.Debug("connected")
	return nil
}

func (service *Scheduler) Stop(ctx context.Context) error {
	if err := service.app.Disconnect(ctx); err != nil {
		return err
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Scheduler) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
