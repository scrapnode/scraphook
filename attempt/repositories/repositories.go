package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Repo struct {
	Database database.Database
	Message  MessageRepo
	Request  RequestRepo
	Response ResponseRepo
	Endpoint EndpointRepo
}

type MessageRepo interface {
	Put(msg *entities.Message) error
}

type RequestRepo interface {
	Put(msg *entities.Request) error
}

type ResponseRepo interface {
	Put(msg *entities.Response) error
}

type EndpointScanResult struct {
	database.ScanResult
	Records []entities.Endpoint
}
type EndpointRepo interface {
	Scan(query *database.ScanQuery) (*EndpointScanResult, error)
}
