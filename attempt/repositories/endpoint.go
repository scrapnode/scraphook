package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Endpoint interface {
	Scan(query *database.ScanQuery) (*EndpointScanResult, error)
}

type EndpointScanResult struct {
	database.ScanResult
	Records []entities.Endpoint
}
