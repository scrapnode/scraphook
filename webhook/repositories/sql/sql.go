package sql

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/database/sql"
	"github.com/scrapnode/scraphook/webhook/repositories"
)

func New(ctx context.Context, cfg *database.Configs) (*repositories.Repo, error) {
	db, err := sql.New(ctx, cfg)
	if err != nil {
		return nil, err
	}

	repo := &repositories.Repo{
		Database: db,
		Webhook:  &WebhookRepo{db: db, clock: clock.New()},
	}
	return repo, nil
}
