package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Webhook interface {
	Save(webhook *entities.Webhook) error
	VerifyOwnership(workspaceId, webhookId string) (bool, error)
	Get(workspaceId, webhookId string) (*entities.Webhook, error)
	List(query *WebhookListQuery) (*WebhookListResult, error)
	Delete(workspaceId, webhookId string) error
}

type WebhookListQuery struct {
	database.ScanQuery
	WorkspaceId string
}
type WebhookListResult struct {
	database.ScanResult
	Data []entities.Webhook
}
