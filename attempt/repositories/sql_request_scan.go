package repositories

import (
	"fmt"
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
	"gorm.io/gorm"
)

func (repo *SqlRequest) Scan(query *RequestScanQuery) (*RequestScanResult, error) {
	conn := repo.db.GetConn().(*gorm.DB)
	tx := conn.Model(&entities.Request{}).
		Where("status in ?", []int{entities.REQ_STATUS_INIT, entities.REQ_STATUS_ATTEMPT}).
		Limit(query.Size).
		Order("id ASC")
	if len(query.Filters) > 0 {
		for col, value := range query.Filters {
			tx = tx.Where(fmt.Sprintf("%s = ?", col), value)
		}
	}
	if query.After > 0 {
		tx = tx.Where("timestamps >= ?", query.After)
	}
	if query.Before > 0 {
		tx = tx.Where("timestamps <= ?", query.Before)
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
	if len(requests) < query.Size {
		return &RequestScanResult{Records: requests}, nil
	}

	// because we order by id with direction ASC
	// so the last item hold the cursor
	results := &RequestScanResult{
		ScanResult: database.ScanResult{Cursor: requests[len(requests)-1].Id},
		Records:    requests,
	}
	return results, nil
}
