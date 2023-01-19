package trigger

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/scrapnode/scrapcore/transport"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/attempt/application"
	"go.uber.org/zap"
)

func New(ctx context.Context, app *application.App) transport.Transport {
	logger := xlogger.FromContext(ctx).With("service", "trigger")
	return &Trigger{app: app, logger: logger}
}

type Trigger struct {
	app    *application.App
	logger *zap.SugaredLogger
	cron   *cron.Cron
	jobs   []cron.EntryID
}

func (service *Trigger) Start(ctx context.Context) error {
	if err := service.app.Connect(ctx); err != nil {
		return err
	}

	service.cron = cron.New()

	id, err := service.cron.AddFunc(service.app.Configs.Trigger.CronPattern, UseTriggerRequest(service.app))
	if err != nil {
		return err
	}
	service.jobs = append(service.jobs, id)

	service.cron.Start()
	service.logger.Debug("connected")
	return nil
}

func (service *Trigger) Stop(ctx context.Context) error {
	cronctx := service.cron.Stop()
	// wait for all tasks are done
	<-cronctx.Done()

	if err := service.app.Disconnect(ctx); err != nil {
		service.logger.Errorw("disconnect app got error", "error", err.Error())
	}

	service.logger.Debug("disconnected")
	return nil
}

func (service *Trigger) Run(ctx context.Context) error {
	// with debug mode, we want to run the task immediately
	if service.app.Configs.Debug() {
		for _, id := range service.jobs {
			service.cron.Entry(id).Job.Run()
		}
	}

	service.logger.Debugw("running", "cron_pattern", service.app.Configs.Trigger.CronPattern)
	service.cron.Run()
	return nil
}
