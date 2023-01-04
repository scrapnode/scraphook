package entities

type Endpoint struct {
	WorkspaceId string `json:"workspace_id" gorm:"index:ws_wh,priority:30"`
	WebhookId   string `json:"webhook_id" gorm:"index:ws_wh,priority:20"`
	Id          string `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"size:256"`
	// http, gRPC
	Uri string `json:"uri" gorm:"size:1024"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`

	Rules []EndpointRule
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
}

func (endpointRule *EndpointRule) TableName() string {
	return "endpoint_rules"
}
