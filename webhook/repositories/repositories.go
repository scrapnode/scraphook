package repositories

import (
	"context"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

func New(ctx context.Context, cfg *database.Configs) (*Repo, error) {
	// Parse the DSN and init corresponding repository
	// Example: uri, err := url.Parse(cfg.Dsn)
	return NewSql(ctx, cfg)
}

type Repo struct {
	Database database.Database
	Webhook  Webhook
}

type Webhook interface {
	GetToken(webhookId, token string) (*entities.WebhookToken, error)
	GetEndpoints(webhookId string) ([]*entities.Endpoint, error)
}
