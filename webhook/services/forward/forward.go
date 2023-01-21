package forward

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/events"
	"github.com/scrapnode/scraphook/webhook/application"
	"go.uber.org/zap"
)

type Forward struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup func() error
}

func New(ctx context.Context, app *application.App) *Forward {
	logger := xlogger.FromContext(ctx).With("service", "examiner")
	return &Forward{app: app, logger: logger}
}

func (service *Forward) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.SCHEDULE_REQUEST}
	queue := service.app.Configs.MsgBus.QueueName
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, queue, UseSubscriber(service.app))
	if err != nil {
		return err
	}

	service.cleanup = cleanup
	service.logger.Debugw("connected", "queue_name", queue)
	return nil
}

func (service *Forward) Stop(ctx context.Context) error {
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

func (service *Forward) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}