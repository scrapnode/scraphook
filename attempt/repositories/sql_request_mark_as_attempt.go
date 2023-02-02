package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlRequest) MarkAsAttempt(ids []string) error {
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Model(&entities.Request{}).
		Where("id in ?", ids).
		Update("status", entities.REQ_STATUS_ATTEMPT)
	return tx.Error
}
