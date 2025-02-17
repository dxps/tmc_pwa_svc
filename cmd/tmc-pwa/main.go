package main

import (
	"log"
	"net/http"

	"github.com/dxps/tmc-pwa/internal/ui/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {

	app.Route("/about", func() app.Composer { return &pages.About{} })
	app.Route("/", func() app.Composer { return &pages.Homepage{} })

	app.RunWhenOnBrowser()

	appHandler := &app.Handler{
		Name:            "TM Community",
		ShortName:       "TMC",
		Description:     "TM Community solution",
		Title:           "TMC Community",
		LoadingLabel:    " ",
		Icon:            app.Icon{Default: "/web/loading.png", SVG: "/web/favicon.svg"},
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
		Styles:          []string{"/web/main.css"},
	}

	s := http.Server{
		Addr:    ":8000",
		Handler: appHandler,
	}

	log.Println("Listening on http://localhost:8000 ...")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
