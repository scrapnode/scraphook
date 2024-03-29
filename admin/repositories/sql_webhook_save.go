package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlWebhook) Save(webhook *entities.Webhook) error {
	updates := map[string]interface{}{
		"name":       webhook.Name,
		"updated_at": webhook.UpdatedAt,
	}
	clauses := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(updates),
	}
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Clauses(clauses).Create(webhook)
	return tx.Error
}
