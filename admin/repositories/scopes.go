package repositories

import "gorm.io/gorm"

func UseWorkspaceScope(ws string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("workspace_id = ?", ws)
	}
}
