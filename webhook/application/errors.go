package application

import "errors"

var (
	ErrWebhookNotFound     = errors.New("webserver: webhook is not found")
	ErrWebhookTokenInvalid = errors.New("webserver: webhook token is not valid")
	ErrEventDataInvalid    = errors.New("scheduler: event data is not valid")
	ErrGetEndpointsFail    = errors.New("scheduler.request: could not get endpoints of webhook")
	ErrNoEndpoints         = errors.New("scheduler.request: no endpoint was found")
)
