package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Get(workspaceId, webhookId, endpointId string) (*entities.Endpoint, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var endpoint entities.Endpoint
	tx := conn.Model(&entities.Endpoint{}).
		Scopes(UseWorkspaceScope(workspaceId)).
		Scopes(UseWebhookScope(webhookId)).
		Where("id = ?", endpointId).
		First(&endpoint)
	return &endpoint, tx.Error
}
