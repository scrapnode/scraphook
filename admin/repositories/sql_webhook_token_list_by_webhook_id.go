package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) ListByWebhookId(webhookId string) ([]entities.WebhookToken, error) {
	conn := repo.db.GetConn().(*gorm.DB)
	var tokens []entities.WebhookToken
	tx := conn.Model(&entities.WebhookToken{}).
		Where("webhook_id = ?", webhookId).
		Find(&tokens)
	return tokens, tx.Error
}
