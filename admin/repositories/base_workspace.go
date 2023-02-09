package repositories

import "github.com/scrapnode/scraphook/entities"

type Workspace interface {
	Get(id string) (*entities.Workspace, error)
}
