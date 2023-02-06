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
		Database:     db,
		Workspace:    &SqlWorkspace{db: db},
		Webhook:      &SqlWebhook{db: db},
		WebhookToken: &SqlWebhookToken{db: db},
		Endpoint:     &SqlEndpoint{db: db},
	}
	return repo, nil
}
