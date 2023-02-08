package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type RequestTrigger struct {
	Bucket string `json:"bucket"`
	Start  int64  `json:"start"`
	End    int64  `json:"end"`

	WebhookId  string `json:"webhook_id"`
	EndpointId string `json:"endpoint_id"`
	Id         string `json:"id"`
}

func (trigger *RequestTrigger) UseId() {
	trigger.Id = utils.NewId("ep")
}

func (trigger *RequestTrigger) Key() string {
	keys := []string{
		trigger.WebhookId,
		trigger.EndpointId,
		trigger.Id,
	}
	return strings.Join(keys, "/")
}
