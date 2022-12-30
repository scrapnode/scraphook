package repositories

import (
	"context"
	databaseconfigs "github.com/scrapnode/scrapcore/database/configs"
	databasesql "github.com/scrapnode/scrapcore/database/sql"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"github.com/scrapnode/scraphook/webhook/repositories/sql"
)

func New(ctx context.Context, cfg *databaseconfigs.Configs) (*interfaces.Repo, error) {
	// by default we will use SQL database
	// if you want to use another database, use init it here
	db, err := databasesql.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return sql.New(ctx, db), err
}
