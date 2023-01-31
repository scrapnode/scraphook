package application

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scrapcore/auth"
	"github.com/scrapnode/scrapcore/msgbus"
	"github.com/scrapnode/scrapcore/xcache"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scrapcore/xmonitor"
	"github.com/scrapnode/scraphook/admin/configs"
	"github.com/scrapnode/scraphook/admin/repositories"
	"go.uber.org/zap"
	"sync"
)

func New(ctx context.Context, cfg *configs.Configs) (*App, error) {
	app := &App{
		Configs: cfg,
		Logger:  xlogger.FromContext(ctx).With("pkg", "scraphook.attempt.application"),
		Clock:   clock.New(),
		Root:    auth.NewAccessKey(cfg.Root.AccessKeyId, cfg.Root.AccessKeySecret),
	}

	monitor, err := xmonitor.New(ctx, app.Configs.Monitor)
	if err != nil {
		return nil, err
	}
	app.Monitor = monitor
	// share monitor across services via context
	ctx = xmonitor.WithContext(ctx, app.Monitor)

	bus, err := msgbus.New(ctx, app.Configs.MsgBus)
	if err != nil {
		return nil, err
	}
	app.MsgBus = bus

	cache, err := xcache.New(ctx, cfg.Cache)
	if err != nil {
		return nil, err
	}
	app.Cache = cache

	repo, err := repositories.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}
	app.Repo = repo

	return app, nil
}

type App struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Clock   clock.Clock

	// services
	Root    *auth.AccessKey
	MsgBus  msgbus.MsgBus
	Monitor xmonitor.Monitor
	Cache   xcache.Cache
	Repo    *repositories.Repo

	mu sync.Mutex
}

func (app *App) Connect(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	if err := app.Root.Connect(ctx); err != nil {
		return err
	}
	if err := app.Cache.Connect(ctx); err != nil {
		return err
	}
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

	if app.Cache != nil {
		if err := app.Cache.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	if app.Root != nil {
		if err := app.Root.Disconnect(ctx); err != nil {
			app.Logger.Error(err)
		}
	}

	app.Logger.Debug("disconnected")
	return nil
}
