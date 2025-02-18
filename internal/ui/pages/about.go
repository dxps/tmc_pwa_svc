package pages

import (
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type About struct {
	app.Compo
}

func (h *About) Render() app.UI {

	return app.Div().Class(
		"flex flex-col min-h-screen bg-gray-100",
	).Body(
		&comps.Navbar{},
		app.Div().
			Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").
			Body(
				app.H1().Text("About"),
				app.A().Href("/").Text("Back to home"),
			),
	)
}
