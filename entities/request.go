package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
	"time"
)

var (
	REQ_STATUS_ATTEMPT = -1
	REQ_STATUS_INIT    = 0
	REQ_STATUS_DONE    = 1
)

type Request struct {
	Timestamps int64  `json:"timestamps"`
	Bucket     string `json:"bucket"`

	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
	Id          string `json:"id"`

	// http, gRPC
	Uri     string `json:"uri"`
	Status  int    `json:"status"`
	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method"`
}

func (request *Request) TableName() string {
	return "requests"
}

func (request *Request) UseId() {
	request.Id = utils.NewId("req")
}

func (request *Request) UseTs(tpl string, t time.Time) {
	request.Bucket, request.Timestamps = utils.NewBucket(tpl, t)
}

func (request *Request) Key() string {
	keys := []string{
		request.WorkspaceId,
		request.WebhookId,
		request.EndpointId,
		request.Id,
	}
	return strings.Join(keys, "/")
}
