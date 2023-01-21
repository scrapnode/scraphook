package repositories

type MessageRepo interface {
	Put(msg *entities.Message) error
}
