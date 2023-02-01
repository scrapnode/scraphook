package repositories

import "github.com/scrapnode/scraphook/entities"

type Webhook interface {
	Save(webhook *entities.Webhook) error
	BelongToWorkspace(workspaceId, webhookId string) (bool, error)
	Get(workspaceId, webhookId string) (*entities.Webhook, error)
}
