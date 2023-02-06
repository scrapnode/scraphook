package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewWebhookServer(app *application.App) *WebhookServer {
	return &WebhookServer{
		app:    app,
		save:   application.NewWebhookSave(app),
		list:   application.NewWebhookList(app),
		get:    application.NewWebhookGet(app),
		delete: application.NewWebhookDelete(app),
	}
}

type WebhookServer struct {
	protos.WebhookServer
	app    *application.App
	save   pipeline.Pipe
	list   pipeline.Pipe
	get    pipeline.Pipe
	delete pipeline.Pipe
}
