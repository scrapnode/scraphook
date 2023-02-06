package protos

import (
	"github.com/samber/lo"
	"github.com/scrapnode/scraphook/entities"
)

func ConvertWebhookToRecord(webhook *entities.Webhook, tokens []entities.WebhookToken) *WebhookRecord {
	record := &WebhookRecord{
		WorkspaceId: webhook.WorkspaceId,
		Id:          webhook.Id,
		Name:        webhook.Name,
		CreatedAt:   webhook.CreatedAt,
		UpdatedAt:   webhook.UpdatedAt,
	}
	if tokens != nil && len(tokens) > 0 {
		record.Tokens = lo.Map(tokens, func(item entities.WebhookToken, _ int) *WebhookTokenRecord {
			return &WebhookTokenRecord{
				WebhookId: item.WebhookId,
				Id:        item.Id,
				Token:     item.Token,
				CreatedAt: item.CreatedAt,
			}
		})
	}
	return record
}
