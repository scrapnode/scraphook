package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlEndpointRule) Save(rule *entities.EndpointRule) error {
	updates := map[string]interface{}{
		"rule":       rule.Rule,
		"negative":   rule.Negative,
		"priority":   rule.Priority,
		"updated_at": rule.UpdatedAt,
	}
	clauses := clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(updates),
	}
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Clauses(clauses).Create(rule)
	return tx.Error
}
