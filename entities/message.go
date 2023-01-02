package entities

import "github.com/scrapnode/scrapcore/utils"

type Message struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Timestamps int64  `json:"timestamps" gorm:"autoCreateTime:milli"`

	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method" gorm:"size:64"`
}

func (msg *Message) WithId() bool {
	// only set data if it wasn't set yet
	if msg.Id != "" {
		return false
	}

	msg.Id = utils.NewId("msg")
	return true
}
