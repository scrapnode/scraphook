package repositories

import (
	"errors"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWebhook) GetToken(webhookId, token string) (*entities.WebhookToken, error) {
	whtoken := &entities.WebhookToken{}

	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.
		Model(&entities.WebhookToken{}).
		Preload("Webhook").
		Where("webhook_id = ? AND token = ?", webhookId, token).
		First(whtoken)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, database.ErrRecordNotFound
		}

		return nil, database.ErrQueryFailed
	}

	return whtoken, nil
}
