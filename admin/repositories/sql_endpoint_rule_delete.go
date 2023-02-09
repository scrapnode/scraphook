package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpointRule) Delete(endpointId, ruleId string) error {
	conn := repo.db.Conn().(*gorm.DB)

	rule := &entities.EndpointRule{Id: ruleId}
	tx := conn.
		Scopes(UseEndpointScope(rule, endpointId)).
		Delete(rule)
	return tx.Error
}
