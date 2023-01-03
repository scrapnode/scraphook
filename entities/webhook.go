package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Webhook struct {
	WorkspaceId string `json:"workspace_id" gorm:"index"`
	Id          string `json:"id" gorm:"primaryKey"`

	Name string `json:"name" gorm:"size:256"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (wh *Webhook) TableName() string {
	return "webhooks"
}

func (wh *Webhook) WithId() bool {
	// only set data if it wasn't set yet
	if wh.Id != "" {
		return false
	}

	wh.Id = utils.NewId("wh")
	return true
}

func (wh *Webhook) Key() string {
	keys := []string{
		wh.WorkspaceId,
		wh.Id,
	}
	return strings.Join(keys, "/")
}

type WebhookToken struct {
	WebhookId string `json:"webhook_id" gorm:"index:ws_wh,priority:20"`
	Id        string `json:"id" gorm:"primaryKey"`

	Token string `json:"token" gorm:"<-:create,unique,not null,size:256"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	DeletedAt int64 `json:"deleted_at" gorm:"default:0"`

	Webhook *Webhook
}

func (wht *WebhookToken) TableName() string {
	return "webhook_tokens"
}

func (wht *WebhookToken) Censor() {
	wht.Token = utils.Censor(wht.Token, 5)
}

func (wht *WebhookToken) WithId() bool {
	// only set data if it wasn't set yet
	if wht.Id != "" {
		return false
	}

	wht.Id = utils.NewId("wht")
	return true
}
