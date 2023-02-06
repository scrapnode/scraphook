package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewEndpointServer(app *application.App) *EndpointServer {
	return &EndpointServer{
		app:    app,
		save:   application.NewEndpointSave(app),
		list:   application.NewEndpointList(app),
		get:    application.NewEndpointGet(app),
		delete: application.NewEndpointDelete(app),
	}
}

type EndpointServer struct {
	protos.EndpointServer
	app    *application.App
	save   pipeline.Pipe
	list   pipeline.Pipe
	get    pipeline.Pipe
	delete pipeline.Pipe
}
