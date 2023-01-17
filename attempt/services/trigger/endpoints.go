package trigger

import "github.com/scrapnode/scraphook/attempt/application"

func UseEndpoints(app *application.App) func() {
	return func() {
		app.Logger.Debugw("running.....................................................................")
	}
}
