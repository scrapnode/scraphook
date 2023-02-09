package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Delete(webhookId, tokenId string) error {
	conn := repo.db.Conn().(*gorm.DB)

	model := &entities.WebhookToken{Id: tokenId}
	tx := conn.
		Scopes(UseWebhookScope(model, webhookId)).
		Delete(model)
	return tx.Error
}
