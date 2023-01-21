package repositories

import "github.com/scrapnode/scraphook/entities"

type ResponseRepo interface {
	Put(msg *entities.Response) error
}
