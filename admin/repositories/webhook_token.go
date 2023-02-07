package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type WebhookToken interface {
	Create(tokens *[]entities.WebhookToken) error
	Get(webhookId, tokenId string) (*entities.WebhookToken, error)
	List(query *WebhookTokenListQuery) (*WebhookTokenListResult, error)
	Delete(webhookId, tokenId string) error
}

type WebhookTokenListQuery struct {
	database.ScanQuery
	WebhookId string
}
type WebhookTokenListResult struct {
	database.ScanResult
	Data []entities.WebhookToken
}
