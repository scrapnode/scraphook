package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Message struct {
	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id" `
	Id          string `json:"id"`
	Timestamps  int64  `json:"timestamps"`

	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method"`
}

func (msg *Message) UseId() {
	msg.Id = utils.NewId("msg")
}

func (msg *Message) Key() string {
	keys := []string{
		msg.WorkspaceId,
		msg.WebhookId,
		msg.Id,
	}
	return strings.Join(keys, "/")
}
