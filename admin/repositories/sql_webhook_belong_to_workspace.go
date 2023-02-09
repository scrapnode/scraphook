package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) VerifyOwnership(workspaceId, webhookId string) (bool, error) {
	var webhook entities.Webhook
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Model(&entities.Webhook{}).
		Scopes(UseWorkspaceScope(workspaceId)).
		Where("id = ?", webhookId).First(&webhook)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
