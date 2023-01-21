package sql

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/attempt/repositories"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *RequestRepo) Scan(query *repositories.RequestScanQuery) (*repositories.RequestScanResult, error) {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Model(&entities.Request{}).
		Where("status in ?", []int{entities.REQ_STATUS_INIT, entities.REQ_STATUS_ATTEMPT}).
		Limit(query.Limit).
		Order("id ASC")
	if query.Before > 0 {
		tx = tx.Where("timestamps <= ?", query.Before)
	}
	if query.After > 0 {
		tx = tx.Where("timestamps >= ?", query.After)
	}
	if query.Cursor != "" {
		tx = tx.Where("id > ?", query.Cursor)
	}

	var requests []entities.Request
	if tx.Scan(&requests); tx.Error != nil {
		return nil, tx.Error
	}
	// total records we got less than the number we requested
	// that mean in database, we have no more record to fetch
	if len(requests) < query.Limit {
		return &repositories.RequestScanResult{Records: requests}, nil
	}

	// because we order by id with direction ASC
	// so the last item hold the cursor
	results := &repositories.RequestScanResult{
		ScanResult: database.ScanResult{Cursor: requests[len(requests)-1].Id},
		Records:    requests,
	}
	return results, nil
}
