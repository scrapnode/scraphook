package trigger

import "github.com/scrapnode/scraphook/attempt/application"

func UseMessage(app *application.App) func() {
	return func() {
		app.Logger.Debugw("running.....................................................................")
	}
}
