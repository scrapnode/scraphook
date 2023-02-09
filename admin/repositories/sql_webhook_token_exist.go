package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Exist(workspaceId, webhookId, tokenId string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	model := &entities.WebhookToken{}
	tx := conn.Model(model).
		Joins("LEFT JOIN webhooks ON webhooks.id = webhook_tokens.webhook_id").
		Scopes(UseWorkspaceScope(&entities.Webhook{}, workspaceId)).
		Scopes(UseWebhookScope(model, webhookId)).
		Where("webhook_tokens.id = ?", tokenId).
		Count(&count)
	return count > 0, tx.Error
}
