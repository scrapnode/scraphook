package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Delete(webhookId, endpointId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(tx *gorm.DB) error {
		endpointRuleTx := tx.
			Scopes(UseEndpointScope(endpointId)).
			Delete(&entities.EndpointRule{})
		if endpointRuleTx.Error != nil {
			return endpointRuleTx.Error
		}

		endpointTx := tx.
			Scopes(UseWebhookScope(webhookId)).
			Delete(&entities.Endpoint{Id: endpointId})
		if endpointTx.Error != nil {
			return endpointTx.Error
		}

		return nil
	})
}
