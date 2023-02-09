package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Exist(workspaceId, webhookId, endpointId string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	model := &entities.Endpoint{}
	tx := conn.Model(model).
		Joins("LEFT JOIN webhooks ON webhooks.id = endpoints.webhook_id").
		Scopes(UseWorkspaceScope(&entities.Webhook{}, workspaceId)).
		Scopes(UseWebhookScope(model, webhookId)).
		Where("endpoints.id = ?", endpointId).
		Count(&count)
	return count > 0, tx.Error
}
