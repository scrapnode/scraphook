package application

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	databasesql "github.com/scrapnode/scrapcore/database/sql"
	"github.com/scrapnode/scrapcore/msgbus"
	msgbusnats "github.com/scrapnode/scrapcore/msgbus/nats"
	"github.com/scrapnode/scrapcore/xlogger"
	"github.com/scrapnode/scraphook/webhook/configs"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"github.com/scrapnode/scraphook/webhook/repositories/sql"
	"go.uber.org/zap"
	"sync"
)

func New(ctx context.Context, cfg *configs.Configs) (*App, error) {
	app := &App{Configs: cfg, Logger: xlogger.FromContext(ctx).With("pkg", "scraphook.webhook.application")}

	// use SQL database by default
	db, err := databasesql.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}
	app.Database = db
	app.Repo = sql.New(ctx, db)

	// use Nats msgbus by default
	bus, err := msgbusnats.New(ctx, app.Configs.MsgBus)
	if err != nil {
		return nil, err
	}
	app.MsgBus = bus

	return app, nil
}

type App struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Repo    *interfaces.Repo

	// services
	Database database.Database
	MsgBus   msgbus.MsgBus
	mu       sync.Mutex
}

func (app *App) Connect(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()

	if err := app.Database.Connect(ctx); err != nil {
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

	if app.Database != nil {
		if err := app.Database.Disconnect(ctx); err != nil {
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
