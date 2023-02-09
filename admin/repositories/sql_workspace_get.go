package repositories

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWorkspace) Get(id string) (*entities.Workspace, error) {
	conn := repo.db.Conn().(*gorm.DB)
	workspace := &entities.Workspace{}
	tx := conn.Model(workspace).
		Where("id = ?", id).
		First(workspace)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return workspace, nil
}
