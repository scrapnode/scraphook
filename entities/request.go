package entities

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

var (
	REQ_STATUS_ATTEMPT = -1
	REQ_STATUS_INIT    = 0
	REQ_STATUS_DONE    = 1
)

type Request struct {
	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
	// the request & response id must be the same value
	// so we can cooperate them
	Id string `json:"id"`

	// http, gRPC
	Uri     string              `json:"uri"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
	Method  string              `json:"method"`

	Timestamps int64 `json:"timestamps"`
}

func (request *Request) WithId() bool {
	// only set data if it wasn't set yet
	if request.Id != "" {
		return false
	}

	request.Id = utils.NewId("req")
	return true
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

func (request *Request) Marshal() []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, _ := json.Marshal(&request)
	return bytes
}

func (request *Request) Unmarshal(data string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(data), request)
}
