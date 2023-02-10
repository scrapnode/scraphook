package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) Get(workspaceId, id string) (*entities.Webhook, error) {
	conn := repo.db.Conn().(*gorm.DB)

	webhook := &entities.Webhook{}
	tx := conn.Model(webhook).
		Scopes(UseWorkspaceScope(webhook, workspaceId)).
		Where("id = ?", id).
		First(webhook)
	return webhook, tx.Error
}
