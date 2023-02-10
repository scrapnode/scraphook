package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Webhook interface {
	Save(webhook *entities.Webhook) error
	VerifyExisting(workspaceId, id string) (bool, error)
	Get(workspaceId, id string) (*entities.Webhook, error)
	List(query *WebhookListQuery) (*WebhookListResult, error)
	Delete(workspaceId, id string) error
}

type WebhookListQuery struct {
	database.ScanQuery
	WorkspaceId string
}
type WebhookListResult struct {
	database.ScanResult
	Data []entities.Webhook
}
