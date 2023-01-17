package message

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/capture/application"
	"github.com/scrapnode/scraphook/events"
	"go.uber.org/zap"
)

func New(ctx context.Context, app *application.App) transport.Transport {
	logger := xlogger.FromContext(ctx).With("service", "message")
	return &Message{app: app, logger: logger}
}

type Message struct {
	app    *application.App
	logger *zap.SugaredLogger

	cleanup func() error
}

func (service *Message) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	sample := &msgbus.Event{Workspace: "*", App: "*", Type: events.MESSAGE}
	queue := service.app.Configs.MsgBus.QueueName
	cleanup, err := service.app.MsgBus.Sub(ctx, sample, queue, UseSubscriber(service.app))
	if err != nil {
		return err
	}
	service.cleanup = cleanup

	service.logger.Debug("connected")
	return nil
}

func (service *Message) Stop(ctx context.Context) error {
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

func (service *Message) Run(ctx context.Context) error {
	service.logger.Debugw("running")
	return nil
}
