package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewWebhookServer(app *application.App) *WebhookServer {
	return &WebhookServer{
		app:         app,
		create:      application.NewWebhookCreate(app),
		update:      application.NewWebhookUpdate(app),
		list:        application.NewWebhookList(app),
		get:         application.NewWebhookGet(app),
		delete:      application.NewWebhookDelete(app),
		tokenAdd:    application.NewWebhookTokenCreate(app),
		tokenDelete: application.NewWebhookTokenDelete(app),
	}
}

type WebhookServer struct {
	protos.WebhookServer
	app         *application.App
	create      pipeline.Pipe
	update      pipeline.Pipe
	list        pipeline.Pipe
	get         pipeline.Pipe
	delete      pipeline.Pipe
	tokenAdd    pipeline.Pipe
	tokenDelete pipeline.Pipe
}
