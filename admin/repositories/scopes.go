package repositories

import "gorm.io/gorm"

func UseWorkspaceScope(id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("workspace_id = ?", id)
	}
}

func UseWebhookScope(id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("webhook_id = ?", id)
	}
}
