package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type EndpointRule interface {
	Save(endpoint *entities.EndpointRule) error
	Get(endpointId, ruleId string) (*entities.EndpointRule, error)
	List(query *EndpointRuleListQuery) (*EndpointRuleListResult, error)
	Delete(endpointId, ruleId string) error
	Exist(workspaceId, endpointId, ruleId string) (bool, error)
}

type EndpointRuleListQuery struct {
	database.ScanQuery
	EndpointId string
}

type EndpointRuleListResult struct {
	database.ScanResult
	Data []entities.EndpointRule
}
