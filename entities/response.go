package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Timestamps int64  `json:"timestamps"`
	Bucket     string `json:"bucket"`

	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
	RequestId   string `json:"request_id"`
	Id          string `json:"id"`

	Uri     string `json:"uri"`
	Status  int    `json:"status"`
	Headers string `json:"headers"`
	Body    string `json:"body"`
}

func (response *Response) TableName() string {
	return "responses"
}

func (response *Response) UseId() {
	response.Id = utils.NewId("res")
}

func (response *Response) UseTs(tpl string, t time.Time) {
	response.Bucket, response.Timestamps = utils.NewBucket(tpl, t)
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
