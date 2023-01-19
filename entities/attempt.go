package entities

import "strings"

type AttemptTrigger struct {
	Bucket string `json:"bucket"`
	Start  int64  `json:"start"`
	End    int64  `json:"end"`

	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
}

func (trigger *AttemptTrigger) Key() string {
	keys := []string{
		trigger.WorkspaceId,
		trigger.WebhookId,
		trigger.EndpointId,
	}
	return strings.Join(keys, "/")
}
