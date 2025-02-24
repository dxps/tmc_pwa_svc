package attr_def

import (
	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type AttributeDefNewPage struct {
	app.Compo

	apiClient *shttp.ApiClient
}

func NewAttributeDefNewPage(apiClient *shttp.ApiClient) *AttributeDefNewPage {
	return &AttributeDefNewPage{apiClient: apiClient}
}

func (page *AttributeDefNewPage) Render() app.UI {

	return app.Div().Class("flex flex-col min-h-screen bg-gray-100").Body(
		&comps.Navbar{},
		app.Div().Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").Body(
			app.Div().Class("bg-white rounded-lg p-3 min-w-[600px] mt-[min(100px)]").Body(
				app.Div().Class("p-6").Body(
					app.Div().Class("flex justify-between mb-8").Body(
						app.P().Class("text-lg font-medium text-gray-500 antialiased").
							Text("Create an Attribute Definition ..."),
						app.A().Class("text-gray-500 text-xl hover:text-gray-800 px-2 rounded-xl transition duration-200").
							Href("/definitions/attributes").Text("x"),
					),
				),
			),
		),
	)
}
