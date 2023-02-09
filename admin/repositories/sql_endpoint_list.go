package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) List(query *EndpointListQuery) (*EndpointListResult, error) {
	conn := repo.db.Conn().(*gorm.DB)
	model := &entities.Endpoint{}
	tx := conn.Model(model).
		Scopes(UseWebhookScope(model, query.WebhookId)).
		Limit(query.Size).
		Order("id DESC")
	if query.Cursor != "" {
		tx = tx.Where("id < ?", query.Cursor)
	}
	if query.Search != "" {
		filter := "%" + query.Search + "%"
		tx = tx.Where("name LIKE ? OR uri LIKE ?", filter, filter)
	}

	var data []entities.Endpoint
	if tx = tx.Find(&data); tx.Error != nil {
		return nil, tx.Error
	}
	var cursor string
	// if total records less than request size, that mean we have no more data
	// don't set cursor so client knows there is no data anymore
	if len(data) == query.Size {
		cursor = data[len(data)-1].Id
	}

	return &EndpointListResult{database.ScanResult{Cursor: cursor}, data}, tx.Error
}
