package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) BelongToWorkspace(workspaceId, webhookId, endpointId string) (bool, error) {
	var endpoint entities.Endpoint
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Model(&entities.Endpoint{}).
		Scopes(UseWorkspaceScope(workspaceId)).
		Scopes(UseWebhookScope(webhookId)).
		Where("id = ?", endpointId).First(&endpoint)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
