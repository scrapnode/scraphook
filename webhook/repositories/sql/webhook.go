package sql

import (
	"github.com/benbjohnson/clock"
	"gorm.io/gorm"
)

type WebhookRepo struct {
	db    *gorm.DB
	clock clock.Clock
}
