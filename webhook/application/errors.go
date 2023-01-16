package application

import "errors"

var (
	ErrWebhookNotFound     = errors.New("webhook is not found")
	ErrWebhookTokenInvalid = errors.New("webhook token is not valid")
	ErrEventDataInvalid    = errors.New("event data is not valid")
	ErrGetEndpointsFail    = errors.New("could not get endpoints of webhook")
	ErrNoEndpoints         = errors.New("no endpoint was found")
)
