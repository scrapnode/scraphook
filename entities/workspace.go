package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Workspace struct {
	UserId string `json:"user_id"`
	Id     string `json:"id"`

	Name string `json:"name"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

func (ws *Workspace) TableName() string {
	return "workspaces"
}

func (ws *Workspace) Key() string {
	keys := []string{
		ws.UserId,
		ws.Id,
	}
	return strings.Join(keys, "/")
}

func (ws *Workspace) UseId() {
	ws.Id = utils.NewId("ws")
}
