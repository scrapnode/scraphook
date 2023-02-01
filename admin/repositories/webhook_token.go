package repositories

import "github.com/scrapnode/scraphook/entities"

type WebhookToken interface {
	Create(tokens *[]entities.WebhookToken) error
	ListByWebhookId(webhookId string) ([]entities.WebhookToken, error)
}
