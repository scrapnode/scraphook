package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlEndpoint) Scan(query *database.ScanQuery) (*EndpointScanResult, error) {
	conn := repo.db.Conn().(*gorm.DB)
	tx := conn.Model(&entities.Endpoint{}).
		Limit(query.Size).
		Order("id ASC")
	if query.Cursor != "" {
		tx = tx.Where("id > ?", query.Cursor)
	}

	var endpoints []entities.Endpoint
	if tx.Scan(&endpoints); tx.Error != nil {
		return nil, tx.Error
	}
	// total records we got less than the number we requested
	// that mean in database, we have no more record to fetch
	if len(endpoints) < query.Size {
		return &EndpointScanResult{Records: endpoints}, nil
	}

	// because we order by id with direction ASC
	// so the last item hold the cursor
	results := &EndpointScanResult{
		ScanResult: database.ScanResult{Cursor: endpoints[len(endpoints)-1].Id},
		Records:    endpoints,
	}
	return results, nil
}
