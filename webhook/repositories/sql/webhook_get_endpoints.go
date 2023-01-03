package sql

import (
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *WebhookRepo) GetEndpoints(ws, id string) ([]*entities.Endpoint, error) {
	var endpoints []*entities.Endpoint

	conn := repo.db.GetConn().(*gorm.DB)
	rows, err := conn.
		Model(&entities.Endpoint{}).
		Preload("Rules").
		Where("workspace_id =  ? AND webhook_id = ?", ws, id).
		Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var endpoint entities.Endpoint
		// ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
		if err := conn.ScanRows(rows, &endpoint); err != nil {
			return nil, err
		}

		endpoints = append(endpoints, &endpoint)
	}

	return endpoints, err
}
