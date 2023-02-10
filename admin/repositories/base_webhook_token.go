package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type WebhookToken interface {
	Create(token *entities.WebhookToken) error
	Get(webhookId, id string) (*entities.WebhookToken, error)
	List(query *WebhookTokenListQuery) (*WebhookTokenListResult, error)
	Delete(webhookId, id string) error
	Exist(workspaceId, id string) (bool, error)
}

type WebhookTokenListQuery struct {
	database.ScanQuery
	WebhookId string
}
type WebhookTokenListResult struct {
	database.ScanResult
	Data []entities.WebhookToken
}
