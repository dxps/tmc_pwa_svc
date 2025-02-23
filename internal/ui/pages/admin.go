package pages

import (
	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Admin struct {
	app.Compo

	apiClient *shttp.ApiClient
}

func NewAdmin(apiClient *shttp.ApiClient) *Admin {
	return &Admin{apiClient: apiClient}
}

func (a *Admin) Render() app.UI {

	return app.Div().Class(
		"flex flex-col min-h-screen bg-gray-100",
	).Body(
		&comps.Navbar{},
		app.Div().
			Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").
			Body(
				app.Div().Class("bg-white rounded-lg p-3 min-w-[600px] mt-[min(100px)]").
					Body(
						app.Div().Class("p-6").
							Body(
								app.H5().Text("Model Management"),
							),
					),
			),
	)
}
