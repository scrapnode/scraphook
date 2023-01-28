package repositories

import "github.com/scrapnode/scraphook/entities"

type Message interface {
	Put(msg *entities.Message) error
}
