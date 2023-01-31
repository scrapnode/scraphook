package application

import "errors"

var (
	ErrSignFailed    = errors.New("incorrect username or password")
	ErrVerifyFailed  = errors.New("could not verify your identity")
	ErrRefreshFailed = errors.New("could not refresh access token")
)
