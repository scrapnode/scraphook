package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Create(token *entities.WebhookToken) error {
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Create(token)
	return tx.Error
}
