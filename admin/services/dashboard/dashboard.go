package dashboard

import (
	"context"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/admin/application"
	"go.uber.org/zap"
)

func New(ctx context.Context, app *application.App) *Dashboard {
	logger := xlogger.FromContext(ctx).With("service", "dashboard")
	return &Dashboard{app: app, logger: logger}
}

type Dashboard struct {
	app    *application.App
	logger *zap.SugaredLogger
}

func (service *Dashboard) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.logger.Debug("connected")
	return nil
}

func (service *Dashboard) Stop(ctx context.Context) error {

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Dashboard) Run(ctx context.Context) error {
	service.logger.Debugw("running")

	return nil
}
