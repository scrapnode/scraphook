package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Get(workspaceId, webhookId string) (*entities.Webhook, error) {
	conn := repo.db.GetConn().(*gorm.DB)

	var webhook entities.Webhook
	tx := conn.Model(&entities.Webhook{}).
		Scopes(UseWorkspaceScope(workspaceId)).
		Where("id = ?", webhookId).
		First(&webhook)
	return &webhook, tx.Error
}
