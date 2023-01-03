package entities

import "github.com/scrapnode/scrapcore/utils"

type Message struct {
	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id" `
	Id          string `json:"id"`
	Timestamps  int64  `json:"timestamps"`

	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method"`
}

func (msg *Message) WithId() bool {
	// only set data if it wasn't set yet
	if msg.Id != "" {
		return false
	}

	msg.Id = utils.NewId("msg")
	return true
}
