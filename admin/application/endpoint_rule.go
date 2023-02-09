package application

type EndpointRuleReq struct {
	EndpointId string `validate:"required"`
	Id         string `validate:"required"`
}
