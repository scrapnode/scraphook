package entities

import (
	"github.com/scrapnode/scrapcore/utils"
	"strings"
)

type Workspace struct {
	UserId string `json:"user_id" gorm:"index"`
	Id     string `json:"id"  gorm:"primaryKey"`

	Name string `json:"name" gorm:"size:256"`

	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

func (ws *Workspace) TableName() string {
	return "workspaces"
}

func (ws *Workspace) WithId() bool {
	// only set data if it wasn't set yet
	if ws.Id != "" {
		return false
	}

	ws.Id = utils.NewId("ws")
	return true
}

func (ws *Workspace) Key() string {
	keys := []string{
		ws.UserId,
		ws.Id,
	}
	return strings.Join(keys, "/")
}
