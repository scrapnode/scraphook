package application

import "errors"

var (
	ErrEventDataInvalid = errors.New("event data is not valid")
	ErrMessagePutFailed = errors.New("could not put message to database")
)
