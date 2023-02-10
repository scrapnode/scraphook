package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlEndpointRule) Delete(endpointId, id string) error {
	conn := repo.db.Conn().(*gorm.DB)

	rule := &entities.EndpointRule{Id: id}
	tx := conn.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "endpoint_id"}}}).
		Scopes(UseEndpointScope(rule, endpointId)).
		Delete(rule)
	if tx.Error != nil {
		return tx.Error
	}
	// if we deleted sucessfully, returning will assign webhook_id for us
	if rule.EndpointId == "" {
		return errors.New("no endpoint was found")
	}

	return nil
}
