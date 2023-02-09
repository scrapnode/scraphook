package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Delete(webhookId, endpointId string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(tx *gorm.DB) error {
		rule := &entities.EndpointRule{}
		endpointRuleTx := tx.
			Scopes(UseEndpointScope(rule, endpointId)).
			Delete(rule)
		if endpointRuleTx.Error != nil {
			return endpointRuleTx.Error
		}

		endpoint := &entities.Endpoint{Id: endpointId}
		endpointTx := tx.
			Scopes(UseWebhookScope(endpoint, webhookId)).
			Delete(endpoint)
		if endpointTx.Error != nil {
			return endpointTx.Error
		}

		return nil
	})
}
