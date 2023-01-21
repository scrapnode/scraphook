package sql

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (repo *ResponseRepo) Put(res *entities.Response) error {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Clauses(clause.OnConflict{DoNothing: true}).Create(res)
	return tx.Error
}
