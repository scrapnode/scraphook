package sql

import (
	"github.com/benbjohnson/clock"
	"gorm.io/gorm"
)

type WebhookRepo struct {
	conn  *gorm.DB
	clock clock.Clock
}
