package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Exist(webhookId, endpointId string) (bool, error) {
	conn := repo.db.Conn().(*gorm.DB)

	var count int64
	tx := conn.Model(&entities.Endpoint{}).
		Scopes(UseWebhookScope(webhookId)).
		Where("id = ?", endpointId).
		Count(&count)
	return count > 0, tx.Error
}
