package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlRequest) MarkAsDone(id string) error {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Model(&entities.Request{}).
		Where("id = ?", id).
		Update("status", entities.REQ_STATUS_DONE)
	return tx.Error
}
