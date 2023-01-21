package sql

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *RequestRepo) Put(req *entities.Request) error {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Clauses(clause.OnConflict{DoNothing: true}).Create(req)
	return tx.Error
}
