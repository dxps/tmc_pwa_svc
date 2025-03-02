package server

import (
	"fmt"
	"log"
	"net/http"

	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func newAppHandler() *app.Handler {
	return &app.Handler{
		Name:         "TM Community",
		ShortName:    "TMC",
		Description:  "TM Community solution",
		Title:        "TM Community",
		LoadingLabel: " ",
		Icon: app.Icon{
			Default: "/web/images/favicon.svg",
			SVG:     "/web/images/favicon.svg",
		},
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
		Styles:          []string{"/web/styles/main.css"},
	}
}

func InitAndStartWebUiClientSide(uiPort, apiPort int) *http.Server {

	apiClient := shttp.NewApiClient(fmt.Sprintf("http://localhost:%d", apiPort))
	initRoutes(apiClient)

	app.RunWhenOnBrowser()

	uiSrv := http.Server{
		Addr:    fmt.Sprintf(":%d", uiPort),
		Handler: newAppHandler(),
	}

	go func() {
		if err := uiSrv.ListenAndServe(); err != http.ErrServerClosed {
			// TODO: get the startup state from this goroutine and handle it,
			//       instead of doing the fatal exit.
			log.Fatal(err)
		}
	}()

	return &uiSrv
}

func InitAndStartWebUiServerSide(uiPort, apiPort int) *http.Server {

	apiClient := shttp.NewApiClient(fmt.Sprintf("http://localhost:%d", apiPort))
	initRoutes(apiClient)

	app.RunWhenOnBrowser()

	uiSrv := http.Server{
		Addr:    fmt.Sprintf(":%d", uiPort),
		Handler: newCustomHandler(),
	}

	go func() {
		if err := uiSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return &uiSrv
}
