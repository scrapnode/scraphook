package sender

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/application"
	"github.com/scrapnode/scraphook/webhook/configs"
	"go.uber.org/zap"
)

type Sender struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup func() error
}

func New(ctx context.Context, app *application.App) *Sender {
	logger := xlogger.FromContext(ctx).With("service", "scheduler")
	return &Sender{app: app, logger: logger}
}

func (service *Sender) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	// @TODO: change queue name
	sample := &msgbus.Event{Workspace: "*", App: "*", Type: configs.EVENT_TYPE_SCHEDULE_REQ}
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, "sender_sample", UseSubscriber(service.app))
	if err != nil {
		return err
	}

	service.cleanup = cleanup
	service.logger.Debug("connected")
	return nil
}

func (service *Sender) Stop(ctx context.Context) error {
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

func (service *Sender) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
