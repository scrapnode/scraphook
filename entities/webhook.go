package entities

import (
	"fmt"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Webhook struct {
	WorkspaceId string `json:"workspace_id"`
	Id          string `json:"id"`

	Name string `json:"name"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (wh *Webhook) TableName() string {
	return "webhooks"
}

func (wh *Webhook) UseId() {
	wh.Id = utils.NewId("wh")
}

func (wh *Webhook) Key() string {
	keys := []string{
		wh.WorkspaceId,
		wh.Id,
	}
	return strings.Join(keys, "/")
}

type WebhookToken struct {
	WebhookId string `json:"webhook_id"`
	Id        string `json:"id"`

	Token string `json:"token"`

	CreatedAt int64 `json:"created_at"`

	Webhook *Webhook
}

func (wht *WebhookToken) TableName() string {
	return "webhook_tokens"
}

func (wht *WebhookToken) UseId() {
	wht.Id = utils.NewId("wht")
}

func (wht *WebhookToken) UseToken(len int) {
	wht.Token = fmt.Sprintf("wht_%s", utils.RandomString(len))
}

func (wht *WebhookToken) Censor() {
	wht.Token = utils.Censor(wht.Token, 5)
}
