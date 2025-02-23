package pages

import (
	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/dxps/tmc-pwa/internal/shared/model/meta"
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type AttributeDefListPage struct {
	app.Compo

	apiClient *shttp.ApiClient
	entries   []meta.AttributeDef
}

func NewAttributeDefListPage(apiClient *shttp.ApiClient) *AttributeDefListPage {
	return &AttributeDefListPage{apiClient: apiClient}
}

func (page *AttributeDefListPage) Render() app.UI {

	return app.Div().Class("flex flex-col min-h-screen bg-gray-100").Body(
		&comps.Navbar{},
		app.Div().Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").Body(
			app.Div().Class("bg-white rounded-lg p-3 min-w-[600px] mt-[min(100px)]").Body(
				app.Div().Class("p-6").Body(
					app.Div().Class("flex justify-between mb-8").Body(
						app.P().Class("text-lg font-medium text-gray-500 antialiased").
							Text("Attribute Definitions"),
						app.A().Class("text-gray-500 text-3xl font-extralight hover:text-gray-800 px-2 rounded-xl transition duration-200").
							Href("/definitions/attributes/new").Text("+"),
					),
					app.If(len(page.entries) == 0, func() app.UI {
						return app.P().Class("pb-4 text-gray-500").Text("There are no entries.")
					}).Else(func() app.UI {
						return app.Div().Class("flex flex-col").Body(
							app.Div().Class("overflow-x-auto sm:-mx-6 lg:-mx-8").Body(
								app.Table().Class("min-w-full text-left text-gray-500").Body(
									app.TBody().Class("bg-white text-sm").Body(
										app.Range(page.entries).Slice(func(i int) app.UI {
											entry := page.entries[i]
											return app.Tr().Class("border-b").Body(
												app.Td().Class("px-6 py-4 whitespace-nowrap").Text(entry.Name),
												app.Td().Class("px-6 py-4 whitespace-nowrap").Text(entry.Description),
												app.Td().Class("px-6 py-4 whitespace-nowrap").Text(entry.ValueType),
												app.Td().Class("px-6 py-4 whitespace-nowrap").Text(entry.IsRequired),
												app.Td().Class("px-6 py-4 whitespace-nowrap").Body(
													app.A().
														Class("text-gray-500 text-3xl font-extralight hover:text-gray-800 px-2 rounded-xl transition duration-200").
														Href("/definitions/attributes/"+entry.Id).Text("🔗"),
												),
											)
										}),
									),
								),
							),
						)
					}),
				),
			),
		),
	)
}
