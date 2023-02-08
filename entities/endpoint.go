package entities

import "github.com/scrapnode/scrapcore/utils"

type Endpoint struct {
	WebhookId string `json:"webhook_id"`
	Id        string `json:"id"`

	Name string `json:"name"`
	// http, gRPC
	Uri string `json:"uri"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	Rules []EndpointRule
}

func (endpoint *Endpoint) UseId() {
	endpoint.Id = utils.NewId("ep")
}

func (endpoint *Endpoint) TableName() string {
	return "endpoints"
}

type EndpointRule struct {
	EndpointId string `json:"endpoint_id"`
	Id         string `json:"id"`

	Rule     string `json:"rule"`
	Negative bool   `json:"negative"`
	Priority int    `json:"priority"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`

	Endpoint *Endpoint
}

func (rule *EndpointRule) UseId() {
	rule.Id = utils.NewId("epr")
}

func (rule *EndpointRule) TableName() string {
	return "endpoint_rules"
}
