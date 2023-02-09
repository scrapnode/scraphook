package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) VerifyExisting(workspaceId, webhookId string) (bool, error) {
	var webhook entities.Webhook
	conn := repo.db.Conn().(*gorm.DB)
	model := &entities.Webhook{}
	tx := conn.Model(model).
		Scopes(UseWorkspaceScope(model, workspaceId)).
		Where("id = ?", webhookId).
		First(&webhook)
	if tx.Error != nil {
		return false, tx.Error
	}

	return true, nil
}
