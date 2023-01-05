package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"net/http"
	"strings"
)

type Response struct {
	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
	RequestId   string `json:"request_id"`
	Id          string `json:"id"`

	Uri     string `json:"uri"`
	Status  int    `json:"status"`
	Headers string `json:"headers"`
	Body    string `json:"body"`

	Timestamps int64 `json:"timestamps"`
}

func (response *Response) WithId() bool {
	// only set data if it wasn't set yet
	if response.Id != "" {
		return false
	}

	response.Id = utils.NewId("res")
	return true
}

func (response *Response) Key() string {
	keys := []string{
		response.WorkspaceId,
		response.WebhookId,
		response.EndpointId,
		response.RequestId,
		response.Id,
	}
	return strings.Join(keys, "/")
}

func (response *Response) OK() bool {
	return response.Status >= http.StatusOK && response.Status < http.StatusMultipleChoices
}
