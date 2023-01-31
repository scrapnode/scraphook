package dashboard

import (
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

type AccountServer struct {
	protos.AccountServer
	app *application.App
}
