package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Exist(workspaceId, Id string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	tx := conn.Model(&entities.Endpoint{}).
		Joins("LEFT JOIN webhooks ON webhooks.id = endpoints.webhook_id").
		Scopes(UseWorkspaceScope(&entities.Webhook{}, workspaceId)).
		Where("endpoints.id = ?", Id).
		Count(&count)
	return count > 0, tx.Error
}
