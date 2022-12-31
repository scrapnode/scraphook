package interfaces

import "github.com/scrapnode/scraphook/entities"

type Repo struct {
	Webhook WebhookRepo
}

type WebhookRepo interface {
	GetToken(id, token string) (*entities.WebhookToken, error)
}