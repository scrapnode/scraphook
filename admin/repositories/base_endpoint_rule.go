package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type EndpointRule interface {
	Save(rule *entities.EndpointRule) error
	Get(endpointId, id string) (*entities.EndpointRule, error)
	List(query *EndpointRuleListQuery) (*EndpointRuleListResult, error)
	Delete(endpointId, id string) error
	Exist(workspaceId, id string) (bool, error)
}

type EndpointRuleListQuery struct {
	database.ScanQuery
	EndpointId string
}

type EndpointRuleListResult struct {
	database.ScanResult
	Data []entities.EndpointRule
}
