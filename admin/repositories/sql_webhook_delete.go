package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Delete(workspaceId, webhookId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(tx *gorm.DB) error {
		webhookTx := tx.
			Scopes(UseWorkspaceScope(workspaceId)).
			Delete(&entities.Webhook{Id: webhookId})
		if webhookTx.Error != nil {
			return webhookTx.Error
		}

		webhookTokenTx := tx.
			Where("webhook_id = ?", webhookId).
			Delete(&entities.WebhookToken{})
		if webhookTokenTx.Error != nil {
			return webhookTokenTx.Error
		}

		return nil
	})
}
