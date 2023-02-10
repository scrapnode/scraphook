package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlEndpoint) Delete(webhookId, Id string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(db *gorm.DB) error {
		endpoint := &entities.Endpoint{Id: Id}
		tx := db.
			Clauses(clause.Returning{Columns: []clause.Column{{Name: "webhook_id"}}}).
			Scopes(UseWebhookScope(endpoint, webhookId)).
			Delete(endpoint)
		if tx.Error != nil {
			return tx.Error
		}
		// if we deleted sucessfully, returning will assign webhook_id for us
		if endpoint.WebhookId == "" {
			return errors.New("no endpoint was found")
		}

		rule := &entities.EndpointRule{}
		if tx := db.Scopes(UseEndpointScope(rule, Id)).Delete(rule); tx.Error != nil {
			return tx.Error
		}

		return nil
	})
}
