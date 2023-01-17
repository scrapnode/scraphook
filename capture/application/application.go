package application

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/scrapnode/scraphook/capture/configs"
	"github.com/scrapnode/scraphook/capture/repositories"
	"github.com/scrapnode/scraphook/capture/repositories/sql"
	"go.uber.org/zap"
	"sync"
)

func New(ctx context.Context, cfg *configs.Configs) (*App, error) {
	app := &App{
		Configs: cfg,
		Logger:  xlogger.FromContext(ctx).With("pkg", "scraphook.capture.application"),
		Clock:   clock.New(),
	}

	monitor, err := xmonitor.New(ctx, app.Configs.Monitor)
	if err != nil {
		return nil, err
	}
	app.Monitor = monitor
	// share monitor across services via context
	ctx = xmonitor.WithContext(ctx, app.Monitor)

	repo, err := sql.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}
	app.Repo = repo

	bus, err := msgbus.New(ctx, app.Configs.MsgBus)
	if err != nil {
		return nil, err
	}
	app.MsgBus = bus

	return app, nil
}

type App struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger

	// services
	Clock   clock.Clock
	MsgBus  msgbus.MsgBus
	Monitor xmonitor.Monitor
	Repo    *repositories.Repo

	mu sync.Mutex
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
	if err := app.Monitor.Connect(ctx); err != nil {
		return err
	}

	app.Logger.Debug("connected")
	return nil
}

func (app *App) Disconnect(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	if app.Monitor != nil {
		if err := app.Monitor.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	if app.MsgBus != nil {
		if err := app.MsgBus.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	if app.Repo != nil && app.Repo.Database != nil {
		if err := app.Repo.Database.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	app.Logger.Debug("disconnected")
	return nil
}
