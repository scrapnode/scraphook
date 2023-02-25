package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewEndpointRuleServer(app *application.App) *EndpointRuleServer {
	return &EndpointRuleServer{
		app:    app,
		create: application.NewEndpointRuleCreate(app),
		update: application.NewEndpointRuleUpdate(app),
		list:   application.NewEndpointRuleList(app),
		get:    application.NewEndpointRuleGet(app),
		delete: application.NewEndpointRuleDelete(app),
	}
}

type EndpointRuleServer struct {
	protos.EndpointRuleServer
	app    *application.App
	create pipeline.Pipe
	update pipeline.Pipe
	list   pipeline.Pipe
	get    pipeline.Pipe
	delete pipeline.Pipe
}
