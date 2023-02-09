package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlEndpoint) Save(endpoint *entities.Endpoint) error {
	updates := map[string]interface{}{
		"name":       endpoint.Name,
		"uri":        endpoint.Uri,
		"updated_at": endpoint.UpdatedAt,
	}
	clauses := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(updates),
		Where:     clause.Where{Exprs: []clause.Expression{clause.Eq{Column: "webhook_id", Value: endpoint.WebhookId}}},
	}
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Clauses(clauses).Create(endpoint)
	return tx.Error
}
