package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Delete(workspaceId, webhookId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(tx *gorm.DB) error {
		webhookTokenTx := tx.
			Scopes(UseWebhookScope(webhookId)).
			Delete(&entities.WebhookToken{})
		if webhookTokenTx.Error != nil {
			return webhookTokenTx.Error
		}

		webhookTx := tx.
			Scopes(UseWorkspaceScope(workspaceId)).
			Delete(&entities.Webhook{Id: webhookId})
		if webhookTx.Error != nil {
			return webhookTx.Error
		}

		return nil
	})
}
