package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Delete(webhookId, tokenId string) error {
	conn := repo.db.Conn().(*gorm.DB)

	tx := conn.
		Scopes(UseWorkspaceScope(webhookId)).
		Delete(&entities.WebhookToken{Id: tokenId})
	return tx.Error
}
