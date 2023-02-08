package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpointRule) Delete(endpointId, ruleId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Scopes(UseEndpointScope(endpointId)).
		Delete(&entities.EndpointRule{Id: ruleId})
	return tx.Error
}
