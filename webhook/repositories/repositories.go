package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Repo struct {
	Database database.Database
	Webhook  WebhookRepo
}

type WebhookRepo interface {
	GetToken(id, token string) (*entities.WebhookToken, error)
}
