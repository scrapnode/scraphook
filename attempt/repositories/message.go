package repositories

import "github.com/scrapnode/scraphook/entities"

type MessageRepo interface {
	Put(msg *entities.Message) error
}
