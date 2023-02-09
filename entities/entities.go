package entities

type Model interface {
	TableName() string
	Key() string
	UseId()
}
