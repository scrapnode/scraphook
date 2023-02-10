package repositories

import (
	"errors"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlWebhook) Delete(workspaceId, id string) error {
	conn := repo.db.Conn().(*gorm.DB)
	return conn.Transaction(func(db *gorm.DB) error {

		webhook := &entities.Webhook{Id: id}
		tx := db.
			Clauses(clause.Returning{Columns: []clause.Column{{Name: "workspace_id"}}}).
			Scopes(UseWorkspaceScope(webhook, workspaceId)).
			Delete(webhook)
		if tx.Error != nil {
			return tx.Error
		}
		// if we deleted sucessfully, returning will assign webhook for us
		if webhook.WorkspaceId == "" {
			return errors.New("no webhook was found")
		}

		token := &entities.WebhookToken{}
		if tx := db.Scopes(UseWebhookScope(token, id)).Delete(token); tx.Error != nil {
			return tx.Error
		}
		return nil
	})
}
