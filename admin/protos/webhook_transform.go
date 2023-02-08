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
		CreatedAt:   ConvertMilliToTimestamp(webhook.CreatedAt),
		UpdatedAt:   ConvertMilliToTimestamp(webhook.UpdatedAt),
	}
	if tokens != nil && len(tokens) > 0 {
		record.Tokens = lo.Map(tokens, func(item entities.WebhookToken, _ int) *WebhookTokenRecord {
			return &WebhookTokenRecord{
				WebhookId: item.WebhookId,
				Id:        item.Id,
				Name:      item.Name,
				Token:     item.Token,
				CreatedAt: ConvertMilliToTimestamp(item.CreatedAt),
			}
		})
	}
	return record
}

func ConvertWebhookTokenToRecord(token *entities.WebhookToken) *WebhookTokenRecord {
	record := &WebhookTokenRecord{
		WebhookId: token.WebhookId,
		Id:        token.Id,
		Name:      token.Name,
		Token:     token.Token,
		CreatedAt: ConvertMilliToTimestamp(token.CreatedAt),
	}
	return record
}
