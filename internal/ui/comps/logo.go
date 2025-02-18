package comps

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Logo struct {
	app.Compo
}

func (l *Logo) Render() app.UI {
	return app.Div().Body(
		app.Img().Src("/web/images/logo.png").Alt("logo").Class("h-8"),
	)
}
