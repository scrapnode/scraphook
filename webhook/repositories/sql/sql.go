package sql

import (
	"context"
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/webhook/repositories/interfaces"
	"gorm.io/gorm"
)

func New(ctx context.Context, db database.Database) *interfaces.Repo {
	conn := db.GetConn().(*gorm.DB)

	return &interfaces.Repo{
		Webhook: &WebhookRepo{conn: conn, clock: clock.New()},
	}
}
