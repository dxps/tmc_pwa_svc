package pages

import (
	"encoding/json"
	"log/slog"

	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/dxps/tmc-pwa/internal/shared/model"
	"github.com/dxps/tmc-pwa/internal/ui/comps"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type AboutPage struct {
	app.Compo

	apiClient *shttp.ApiClient
	state     struct {
		Health model.Health
	}
}

func NewAboutPage(apiClient *shttp.ApiClient) *AboutPage {
	return &AboutPage{apiClient: apiClient}
}

func (a *AboutPage) Render() app.UI {

	return app.Div().Class(
		"flex flex-col min-h-screen bg-gray-100",
	).Body(
		&comps.Navbar{},
		app.Div().
			Class("flex flex-col min-h-screen justify-center items-center drop-shadow-2xl").
			Body(
				app.H1().Text("About"),
				app.A().Href("/").Text("Back to home"),
				app.Button().
					Class("bg-gray-400 hover:bg-gray-600 text-white my-8 py-1 px-4 rounded-md").
					Text("Healthcheck").
					OnClick(a.handleHealthcheck),
				app.If(a.state.Health.State != "", func() app.UI {
					return app.Div().Text("Health state: " + a.state.Health.State)
				}),
			),
	)
}

func (a *AboutPage) handleHealthcheck(ctx app.Context, e app.Event) {
	health, err := a.getHealthcheck()
	if err != nil {
		a.state.Health = model.Health{State: "unknown"}
	} else {
		a.state.Health = *health
	}
}

func (a *AboutPage) getHealthcheck() (*model.Health, error) {

	respBody, err := a.apiClient.Get("/health")
	if err != nil {
		slog.Error("getHealthCheck call failed.", "error", err)
		return nil, err
	}
	var health model.Health
	if err := json.Unmarshal(respBody, &health); err != nil {
		slog.Error("getHealthCheck failed to unmarshal json.", "error", err)
		return nil, err
	}
	slog.Info("getHealthCheck call succeeded.", "responseBody", respBody)
	return &health, nil
}
