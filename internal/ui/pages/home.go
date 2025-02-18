package pages

import (
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Homepage struct {
	app.Compo
}

func (h *Homepage) Render() app.UI {
	return app.Div().Class(
		"flex flex-col min-h-screen bg-gray-100",
	).Body(
		&comps.Navbar{},
		app.Div().
			Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").
			Body(
				app.Img().
					Src("/web/images/logo.svg").
					Class("w-[86px] h-[86px] logo_filter"),
				app.H1().Text("TM Community"),
				app.A().Href("/about").Text("About"),
			),
	)
}
