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
	WorkspaceId string `json:"workspace_id" gorm:"index:ws_wh_ep,priority:30"`
	WebhookId   string `json:"webhook_id" gorm:"index:ws_wh_ep,priority:20"`
	EndpointId  string `json:"endpoint_id" gorm:"index:ws_wh_ep,priority:10"`
	Id          string `json:"id" gorm:"primaryKey"`

	Rule     string `json:"rule" gorm:"size:2048"`
	Negative bool   `json:"negative" gorm:"default:false"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (endpointRule *EndpointRule) TableName() string {
	return "endpoint_rules"
}
