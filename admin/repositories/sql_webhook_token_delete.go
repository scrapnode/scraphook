package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Delete(webhookId, Id string) error {
	conn := repo.db.Conn().(*gorm.DB)

	tx := conn.
		Scopes(UseWebhookScope(webhookId)).
		Delete(&entities.WebhookToken{Id: Id})
	return tx.Error
}
