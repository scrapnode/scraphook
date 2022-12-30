package entities

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Response struct {
	WorkspaceId string `json:"workspace_id"`
	WebhookId   string `json:"webhook_id"`
	EndpointId  string `json:"endpoint_id"`
	// the request & response id must be the same value
	// so we can cooperate them
	Id string `json:"id"`

	Uri     string              `json:"uri"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`

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
		response.Id,
	}
	return strings.Join(keys, "/")
}

func (response *Response) Marshal() []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, _ := json.Marshal(&response)
	return bytes
}

func (response *Response) Unmarshal(data string) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal([]byte(data), response)
}
