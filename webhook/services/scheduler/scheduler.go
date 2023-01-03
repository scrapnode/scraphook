package scheduler

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
	"log"
)

type Scheduler struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup func() error
}

func New(ctx context.Context, app *application.App) *Scheduler {
	logger := xlogger.FromContext(ctx).With("service", "scheduler")
	return &Scheduler{app: app, logger: logger}
}

func (service *Scheduler) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	sample := &msgbus.Event{Workspace: "*", App: "*", Type: "*"}
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, "sample", func(event *msgbus.Event) error {
		log.Println(event)
		return nil
	})
	if err != nil {
		return err
	}

	service.cleanup = cleanup
	service.logger.Debug("connected")
	return nil
}

func (service *Scheduler) Stop(ctx context.Context) error {
	if service.cleanup != nil {
		if err := service.cleanup(); err != nil {
			service.logger.Errorw("cleanup subscriber got error", "error", err.Error())
		}
	}

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Scheduler) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
