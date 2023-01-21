package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type RequestScanQuery struct {
	database.ScanQuery
	Bucket string
	Before int64
	After  int64
}

type RequestScanResult struct {
	database.ScanResult
	Records []entities.Request
}

type RequestRepo interface {
	Put(msg *entities.Request) error
	MarkAsDone(id string) error
	Scan(query *RequestScanQuery) (*RequestScanResult, error)
	MarkAsAttempt(ids []string) error
}
