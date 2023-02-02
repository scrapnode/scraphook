package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) BelongToWorkspace(workspaceId, webhookId string) (bool, error) {
	if workspaceId == "" {
		return false, errors.New("workspace id is empty")
	}

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
