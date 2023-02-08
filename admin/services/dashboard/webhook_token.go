package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewWebhookTokenServer(app *application.App) *WebhookTokenServer {
	return &WebhookTokenServer{
		app:    app,
		create: application.NewWebhookTokenCreate(app),
		list:   application.NewWebhookTokenList(app),
		get:    application.NewWebhookTokenGet(app),
		delete: application.NewWebhookTokenDelete(app),
	}
}

type WebhookTokenServer struct {
	protos.WebhookTokenServer
	app    *application.App
	create pipeline.Pipe
	list   pipeline.Pipe
	get    pipeline.Pipe
	delete pipeline.Pipe
}
