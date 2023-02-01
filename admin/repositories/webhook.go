package repositories

import "github.com/scrapnode/scraphook/entities"

type Webhook interface {
	Save(webhook *entities.Webhook) error
	BelongToWorkspace(webhookId, workspaceId string) (bool, error)
}
