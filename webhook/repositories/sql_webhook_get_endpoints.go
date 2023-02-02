package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) GetEndpoints(ws, id string) ([]*entities.Endpoint, error) {
	var endpoints []*entities.Endpoint

	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Model(&entities.Endpoint{}).
		Preload("Rules").
		Where("workspace_id =  ? AND webhook_id = ?", ws, id).
		Find(&endpoints)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return endpoints, nil
}
