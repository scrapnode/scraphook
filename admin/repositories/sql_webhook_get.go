package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Get(workspaceId, webhookId string) (*entities.Webhook, error) {
	conn := repo.db.Conn().(*gorm.DB)

	webhook := &entities.Webhook{}
	tx := conn.Model(webhook).
		Scopes(UseWorkspaceScope(webhook, workspaceId)).
		Where("id = ?", webhookId).
		First(webhook)
	return webhook, tx.Error
}
