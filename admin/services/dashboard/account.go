package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewAccountServer(app *application.App) *AccountServer {
	return &AccountServer{
		app:     app,
		sign:    application.NewAccountSign(app),
		verify:  application.NewAccountVerify(app),
		refresh: application.NewAccountRefresh(app),
	}
}

type AccountServer struct {
	protos.AccountServer
	app     *application.App
	sign    pipeline.Pipe
	verify  pipeline.Pipe
	refresh pipeline.Pipe
}
