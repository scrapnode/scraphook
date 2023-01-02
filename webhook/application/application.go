package application

import (
	"context"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/msgbus/nats"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/repositories"
	"github.com/scrapnode/scraphook/webhook/repositories/sql"
	"go.uber.org/zap"
	"sync"
)

func New(ctx context.Context, cfg *configs.Configs) (*App, error) {
	app := &App{Configs: cfg, Logger: xlogger.FromContext(ctx).With("pkg", "scraphook.webhook.application")}

	repo, err := sql.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}
	app.Repo = repo

	// use Nats msgbus by default
	bus, err := nats.New(ctx, app.Configs.MsgBus)
	if err != nil {
		return nil, err
	}
	app.MsgBus = bus

	return app, nil
}

type App struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Repo    *repositories.Repo

	// services
	MsgBus msgbus.MsgBus
	mu     sync.Mutex
}

func (app *App) Connect(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	if err := app.Repo.Database.Connect(ctx); err != nil {
		return err
	}
	if err := app.MsgBus.Connect(ctx); err != nil {
		return err
	}

	app.Logger.Debug("connected")
	return nil
}

func (app *App) Disconnect(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	if app.Repo != nil && app.Repo.Database != nil {
		if err := app.Repo.Database.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	if app.MsgBus != nil {
		if err := app.MsgBus.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	app.Logger.Debug("disconnected")
	return nil
}
