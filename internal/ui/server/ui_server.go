package server

import (
	"fmt"
	"log"
	"net/http"

	shttp "github.com/dxps/tmc-pwa/internal/shared/http"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func StartWebUiServer(uiPort, apiPort int) *http.Server {

	apiClient := shttp.NewApiClient(fmt.Sprintf("http://localhost:%d", apiPort))
	initRoutes(apiClient)

	app.RunWhenOnBrowser()

	appHandler := &app.Handler{
		Name:         "TM Community",
		ShortName:    "TMC",
		Description:  "TM Community solution",
		Title:        "TM Community",
		LoadingLabel: " ",
		Icon: app.Icon{
			Default: "/web/images/loading.png",
			SVG:     "/web/images/favicon.svg",
		},
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
		Styles:          []string{"/web/styles/main.css"},
	}

	uiSrv := http.Server{
		Addr:    fmt.Sprintf(":%d", uiPort),
		Handler: appHandler,
	}

	go func() {
		if err := uiSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err) // TODO: get the state from this goroutine and handle it.
		}
	}()

	return &uiSrv
}
