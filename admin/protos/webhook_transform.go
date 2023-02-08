package protos

import (
	"github.com/samber/lo"
	"github.com/scrapnode/scraphook/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ConvertWebhookToRecord(webhook *entities.Webhook, tokens []entities.WebhookToken) *WebhookRecord {
	record := &WebhookRecord{
		WorkspaceId: webhook.WorkspaceId,
		Id:          webhook.Id,
		Name:        webhook.Name,
		CreatedAt:   timestamppb.New(time.UnixMilli(webhook.CreatedAt)),
		UpdatedAt:   timestamppb.New(time.UnixMilli(webhook.UpdatedAt)),
	}
	if tokens != nil && len(tokens) > 0 {
		record.Tokens = lo.Map(tokens, func(item entities.WebhookToken, _ int) *WebhookTokenRecord {
			return &WebhookTokenRecord{
				WebhookId: item.WebhookId,
				Id:        item.Id,
				Name:      item.Name,
				Token:     item.Token,
				CreatedAt: timestamppb.New(time.UnixMilli(item.CreatedAt)),
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
		CreatedAt: timestamppb.New(time.UnixMilli(token.CreatedAt)),
	}
	return record
}
