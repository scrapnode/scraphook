package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpointRule) Exist(workspaceId, Id string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	model := &entities.EndpointRule{}
	tx := conn.Model(model).
		Joins("LEFT JOIN endpoints ON endpoints.id = endpoint_rules.endpoint_id").
		Joins("LEFT JOIN webhooks ON webhooks.id = endpoints.webhook_id").
		Scopes(UseWorkspaceScope(&entities.Webhook{}, workspaceId)).
		Where("endpoint_rules.id = ?", Id).
		Count(&count)
	return count > 0, tx.Error
}
