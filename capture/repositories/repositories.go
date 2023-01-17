package repositories

import (
	"github.com/scrapnode/scrapcore/database"
	"github.com/scrapnode/scraphook/entities"
)

type Repo struct {
	Database database.Database
	Message  MessageRepo
}

type MessageRepo interface {
	Put(msg *entities.Message) error
}
