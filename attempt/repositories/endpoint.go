package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type EndpointScanResult struct {
	database.ScanResult
	Records []entities.Endpoint
}

type EndpointRepo interface {
	Scan(query *database.ScanQuery) (*EndpointScanResult, error)
}
