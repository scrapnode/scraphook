package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Endpoint interface {
	Save(endpoint *entities.Endpoint) error
	Exist(workspaceId, webhookId, endpointId string) (bool, error)
	Get(webhookId, endpointId string) (*entities.Endpoint, error)
	List(query *EndpointListQuery) (*EndpointListResult, error)
	Delete(webhookId, endpointId string) error
}

type EndpointListQuery struct {
	database.ScanQuery
	WebhookId string
}
type EndpointListResult struct {
	database.ScanResult
	Data []entities.Endpoint
}
