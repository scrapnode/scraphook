package repositories

import (
	"errors"
	"fmt"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlWorkspace) GetById(id string) (*entities.Workspace, error) {
	if id == "" {
		return nil, errors.New(fmt.Sprintf("workspace #%s is not found", id))
	}

	var workspace entities.Workspace
	conn := repo.db.Conn().(*gorm.DB)
	if tx := conn.Model(&entities.Workspace{}).Where("id = ?", id).First(&workspace); tx.Error != nil {
		return nil, tx.Error
	}

	return &workspace, nil
}
