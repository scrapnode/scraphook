package repositories

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"github.com/scrapnode/scraphook/webhook/repositories/sql"
)

func New(ctx context.Context, db database.Database) (*interfaces.Repo, error) {
	return sql.New(ctx, db), nil
}
