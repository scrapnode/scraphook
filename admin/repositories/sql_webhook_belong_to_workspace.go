package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) BelongToWorkspace(webhookId, workspaceId string) (bool, error) {
	if workspaceId == "" {
		return false, errors.New("workspace id is empty")
	}

	var webhook entities.Webhook
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.
		Model(&entities.Webhook{}).
		Where("id = ? AND workspace_id = ?", webhookId, workspaceId).First(&webhook)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
