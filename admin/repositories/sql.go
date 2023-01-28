package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/admin/repositories"
)

func New(ctx context.Context, cfg *database.Configs) (*repositories.Repo, error) {
	db, err := database.NewSQL(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repo := &repositories.Repo{
		Database: db,
	}
	return repo, nil
}
