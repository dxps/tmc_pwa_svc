package pages

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type About struct {
	app.Compo
}

func (h *About) Render() app.UI {
	return app.Div().Class(
		"flex flex-col min-h-screen bg-gray-100",
	).Body(
		app.H1().Text("About"),
		app.A().Href("/").Text("Back to home"),
	)
}
