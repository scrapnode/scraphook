package repositories

import (
	"fmt"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func UseWorkspaceScope(model entities.Model, id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s.workspace_id = ?", model.TableName()), id)
	}
}

func UseWebhookScope(model entities.Model, id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s.webhook_id = ?", model.TableName()), id)
	}
}

func UseEndpointScope(model entities.Model, id string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s.endpoint_id = ?", model.TableName()), id)
	}
}
