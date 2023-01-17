package sql

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/attempt/repositories"
)

func New(ctx context.Context, cfg *database.Configs) (*repositories.Repo, error) {
	db, err := database.NewSQL(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repo := &repositories.Repo{
		Database: db,
		Message:  &MessageRepo{db: db},
		Request:  &RequestRepo{db: db},
		Response: &ResponseRepo{db: db},
	}
	return repo, nil
}
