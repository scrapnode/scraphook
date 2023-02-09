package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Delete(workspaceId, webhookId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(tx *gorm.DB) error {
		token := &entities.WebhookToken{}
		webhookTokenTx := tx.
			Scopes(UseWebhookScope(token, webhookId)).
			Delete(token)
		if webhookTokenTx.Error != nil {
			return webhookTokenTx.Error
		}

		webhook := &entities.Webhook{Id: webhookId}
		webhookTx := tx.
			Scopes(UseWorkspaceScope(webhook, workspaceId)).
			Delete(webhook)
		if webhookTx.Error != nil {
			return webhookTx.Error
		}

		return nil
	})
}
