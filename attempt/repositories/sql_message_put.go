package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *SqlMessage) Put(msg *entities.Message) error {
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Clauses(clause.OnConflict{DoNothing: true}).Create(msg)
	return tx.Error
}
