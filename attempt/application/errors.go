package application

import "errors"

var (
	ErrEventDataInvalid        = errors.New("event data is not valid")
	ErrMessagePutFailed        = errors.New("could not put message to database")
	ErrRequestPutFailed        = errors.New("could not put request to database")
	ErrResponsePutFailed       = errors.New("could not put response to database")
	ErrMarkRequestAsDoneFailed = errors.New("could not mark request as done")
)
