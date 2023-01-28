package repositories

import "github.com/scrapnode/scraphook/entities"

type Response interface {
	Put(msg *entities.Response) error
}
