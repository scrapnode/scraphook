package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhookToken) Exist(workspaceId, Id string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	tx := conn.Model(&entities.WebhookToken{}).
		Joins("LEFT JOIN webhooks ON webhooks.id = webhook_tokens.webhook_id").
		Scopes(UseWorkspaceScope(&entities.Webhook{}, workspaceId)).
		Where("webhook_tokens.id = ?", Id).
		Count(&count)
	return count > 0, tx.Error
}
