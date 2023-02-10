package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlWebhookToken) Delete(webhookId, id string) error {
	conn := repo.db.Conn().(*gorm.DB)

	token := &entities.WebhookToken{Id: id}
	tx := conn.
		Clauses(clause.Returning{Columns: []clause.Column{{Name: "webhook_id"}}}).
		Scopes(UseWebhookScope(token, webhookId)).
		Delete(token)
	if tx.Error != nil {
		return tx.Error
	}
	// if we deleted sucessfully, returning will assign webhook_id for us
	if token.WebhookId == "" {
		return errors.New("no webhook token was found")
	}

	return nil
}
