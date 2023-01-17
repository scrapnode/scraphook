package sql

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *RequestRepo) Put(msg *entities.Request) error {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Clauses(clause.OnConflict{DoNothing: true}).Create(msg)
	return tx.Error
}
