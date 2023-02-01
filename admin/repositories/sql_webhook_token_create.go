package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Create(tokens *[]entities.WebhookToken) error {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Create(tokens)
	return tx.Error
}
