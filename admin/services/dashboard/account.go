package dashboard

import (
	"github.com/scrapnode/scrapcore/pipeline"
	"github.com/scrapnode/scraphook/admin/application"
	"github.com/scrapnode/scraphook/admin/protos"
)

func NewAccountServer(app *application.App) *AccountServer {
	instrumentName := "admin_account"
	return &AccountServer{
		app:     app,
		sign:    application.NewAccountSign(app, instrumentName),
		verify:  application.NewAccountVerify(app, instrumentName),
		refresh: application.NewAccountRefresh(app, instrumentName),
	}
}

type AccountServer struct {
	protos.AccountServer
	app     *application.App
	sign    pipeline.Pipe
	verify  pipeline.Pipe
	refresh pipeline.Pipe
}
