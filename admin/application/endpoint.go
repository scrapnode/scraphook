package application

type EndpointReq struct {
	WebhookId string `validate:"required"`
	Id        string `validate:"required"`
}
