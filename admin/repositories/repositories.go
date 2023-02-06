package repositories

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
)

func New(ctx context.Context, cfg *database.Configs) (*Repo, error) {
	// Parse the DSN and init corresponding repository
	// Example: uri, err := url.Parse(cfg.Dsn)
	return NewSql(ctx, cfg)
}

type Repo struct {
	Database     database.Database
	Workspace    Workspace
	Webhook      Webhook
	WebhookToken WebhookToken
	Endpoint     Endpoint
}
