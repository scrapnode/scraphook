package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Get(webhookId, Id string) (*entities.WebhookToken, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var token entities.WebhookToken
	tx := conn.Model(&entities.WebhookToken{}).
		Scopes(UseWebhookScope(webhookId)).
		Where("id = ?", Id).
		First(&token)
	return &token, tx.Error
}
