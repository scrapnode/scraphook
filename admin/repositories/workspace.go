package repositories

import "github.com/scrapnode/scraphook/entities"

type Workspace interface {
	GetById(id string) (*entities.Workspace, error)
}
