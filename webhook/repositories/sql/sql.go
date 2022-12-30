package sql

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"gorm.io/gorm"
)

func New(ctx context.Context, db *gorm.DB) *interfaces.Repo {
	return &interfaces.Repo{Webhook: &WebhookRepo{db: db, clock: clock.New()}}
}
