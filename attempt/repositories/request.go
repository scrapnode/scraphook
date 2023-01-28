package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Request interface {
	Put(msg *entities.Request) error
	MarkAsDone(id string) error
	Scan(query *RequestScanQuery) (*RequestScanResult, error)
	MarkAsAttempt(ids []string) error
}

type RequestScanQuery struct {
	database.ScanQuery
	Filters map[string]string
	Bucket  string
	Before  int64
	After   int64
}

type RequestScanResult struct {
	database.ScanResult
	Records []entities.Request
}
