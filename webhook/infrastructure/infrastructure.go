package infrastructure

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

type Infra struct {
	Configs *configs.Configs
	Logger  *zap.SugaredLogger
	Repo    *interfaces.Repo

	// services
	Database database.Database
	MsgBus   msgbus.MsgBus
	mu       sync.Mutex
}

func (infra *Infra) Connect(ctx context.Context) error {
	// always disconnect previous connection before start a new one
	if err := infra.Disconnect(ctx); err != nil {
		return err
	}

	infra.mu.Lock()
	defer infra.mu.Unlock()

	if err := infra.Database.Connect(ctx); err != nil {
		return err
	}
	if err := infra.MsgBus.Connect(ctx); err != nil {
		return err
	}

	infra.Logger.Debug("connected")
	return nil
}

func (infra *Infra) Disconnect(ctx context.Context) error {
	infra.mu.Lock()
	defer infra.mu.Unlock()

	if err := infra.Database.Disconnect(ctx); err != nil {
		infra.Logger.Error(err)
	}
	if err := infra.MsgBus.Disconnect(ctx); err != nil {
		infra.Logger.Error(err)
	}

	infra.Logger.Debug("disconnected")
	return nil
}

func New(ctx context.Context, cfg *configs.Configs) (*Infra, error) {
	infra := &Infra{Configs: cfg, Logger: xlogger.FromContext(ctx).With("service", "scraphook.webhook")}

	// use SQL database by default
	db, err := databasesql.New(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}
	infra.Repo = sql.New(ctx, db.Conn)

	// use Nats msgbus by default
	bus, err := msgbusnats.New(ctx, infra.Configs.MsgBus)
	if err != nil {
		return nil, err
	}
	infra.MsgBus = bus

	return infra, nil
}
