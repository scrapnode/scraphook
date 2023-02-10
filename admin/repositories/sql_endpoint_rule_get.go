package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpointRule) Get(endpointId, id string) (*entities.EndpointRule, error) {
	conn := repo.db.Conn().(*gorm.DB)

	rule := &entities.EndpointRule{}
	tx := conn.Model(rule).
		Scopes(UseEndpointScope(rule, endpointId)).
		Where("id = ?", id).
		First(rule)
	return rule, tx.Error
}
