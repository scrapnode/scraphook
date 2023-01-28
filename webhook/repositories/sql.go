package repositories

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
)

func NewSql(ctx context.Context, cfg *database.Configs) (*Repo, error) {
	db, err := database.NewSQL(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repo := &Repo{
		Database: db,
		Webhook:  &SqlWebhook{db: db},
	}
	return repo, nil
}
