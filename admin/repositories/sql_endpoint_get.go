package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Get(webhookId, id string) (*entities.Endpoint, error) {
	conn := repo.db.Conn().(*gorm.DB)

	endpoint := &entities.Endpoint{}
	tx := conn.Model(endpoint).
		Scopes(UseWebhookScope(endpoint, webhookId)).
		Where("id = ?", id).
		First(endpoint)
	return endpoint, tx.Error
}
