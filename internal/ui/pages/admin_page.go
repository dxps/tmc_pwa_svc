package pages

import (
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type AdminPage struct {
	app.Compo
}

func NewAdminPage() *AdminPage {
	return &AdminPage{}
}

func (a *AdminPage) Render() app.UI {

	return app.Div().Class("flex flex-col min-h-screen bg-gray-100").Body(
		&comps.Navbar{},
		app.Div().
			Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").
			Body(
				app.Div().Class("bg-white rounded-lg p-3 min-w-[600px] mt-[min(100px)]").
					Body(
						app.Div().Class("p-6").
							Body(
								app.H3().Text("Model Management"),
								app.P().Class("block font-sans text-base text-gray-500 leading-relaxed antialiased").
									Text("Manage the definitions and instances of attributes and entities."),
								app.Hr().Class("mt-2 mb-4"),
								app.Div().Class("flex").Body(
									app.Div().Class("pr-3 flex flex-col grow mr-1").Body(
										app.H6().
											Class("px-4 mb-2 pt-2 pb-1 block font-medium text-gray-500 antialiased").
											Text("Definitions"),
										app.A().
											Class("py-2 px-4 rounded-lg transition duration-200").
											Href("/definitions/entities").Text("Entities"),
										app.A().
											Class("py-2 px-4 rounded-lg transition duration-200").
											Href("/definitions/attributes").Text("Attributes"),
									),
									app.Div().Class("pr-3 flex flex-col grow ml-1").Body(
										app.H6().
											Class("px-4 mb-2 pt-2 pb-1 block font-medium text-gray-500 antialiased").
											Text("Instances"),
										app.A().
											Class("py-2 px-4 rounded-lg transition duration-200").
											Href("/instances/entities").Text("Entities"),
										app.A().
											Class("py-2 px-4 rounded-lg transition duration-200").
											Href("/instances/attributes").Text("Attributes"),
									),
								),
							),
					),
			),
	)
}
