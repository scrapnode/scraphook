package sql

import (
	"errors"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scrapcore/database/sql"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *WebhookRepo) GetToken(id, token string) (*entities.WebhookToken, error) {
	whtoken := &entities.WebhookToken{}

	tx := repo.conn.
		Preload("Webhook").
		Scopes(sql.UseNotDeleted(repo.clock)).
		Where("webhook_id = ? AND token = ?", id, token).
		First(whtoken)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, database.ErrRecordNotFound
		}

		return nil, database.ErrQueryFailed
	}

	return whtoken, nil
}
