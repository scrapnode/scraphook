package sql

import (
	"github.com/benbjohnson/clock"
	"github.com/scrapnode/scrapcore/database/sql"
)

type WebhookRepo struct {
	db    *sql.SQL
	clock clock.Clock
}
