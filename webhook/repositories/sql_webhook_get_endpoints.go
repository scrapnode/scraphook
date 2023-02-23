package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) GetEndpoints(webhookId string) ([]*entities.Endpoint, error) {
	var endpoints []*entities.Endpoint

	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Model(&entities.Endpoint{}).
		Preload("Rules", func(db *gorm.DB) *gorm.DB {
			// we have to check negative rule first to ignore them
			// then load rules by their priority
			return db.Order("endpoint_rules.negative DESC, endpoint_rules.priority DESC")
		}).
		Where("webhook_id = ?", webhookId).
		Find(&endpoints)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return endpoints, nil
}
