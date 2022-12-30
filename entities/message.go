package entities

type Message struct {
	Id         string `json:"id" gorm:"primaryKey"`
	Timestamps int64  `json:"timestamps" gorm:"autoCreateTime:milli"`

	Uri     string `json:"uri" gorm:"size:512"`
	Headers string `json:"headers"`
	Body    string `json:"body"`
	Method  string `json:"method" gorm:"size:64"`
}
