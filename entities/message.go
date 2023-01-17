package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
	"time"
)

type Message struct {
	Timestamps int64  `json:"timestamps"`
	Bucket     string `json:"bucket"`

	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id" `
	Id          string `json:"id"`

	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method"`
}

func (msg *Message) UseId() {
	msg.Id = utils.NewId("msg")
}

func (msg *Message) UseTs(tpl string, t time.Time) {
	msg.Bucket, msg.Timestamps = utils.NewBucket(tpl, t)
}

func (msg *Message) Key() string {
	keys := []string{
		msg.WorkspaceId,
		msg.WebhookId,
		msg.Id,
	}
	return strings.Join(keys, "/")
}
