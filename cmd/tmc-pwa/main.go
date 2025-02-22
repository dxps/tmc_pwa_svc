package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"

	"github.com/dxps/tmc-pwa/internal/svc/repos"
	"github.com/dxps/tmc-pwa/internal/ui/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/sethvargo/go-envconfig"
)

func main() {

	var cfg config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		slog.Error("Failed to load config.", "error", err)
		return
	}

	_, err := repos.NewRepos(cfg.Db.Driver, cfg.Db.DSN, 10, 1, "5m")
	if err != nil {
		slog.Error("Failed to connect to database.", "error", err)
	} else {
		slog.Info("Database connection established.")
	}

	app.Route("/about", func() app.Composer { return &pages.About{} })
	app.Route("/", func() app.Composer { return &pages.Homepage{} })

	app.RunWhenOnBrowser()

	appHandler := &app.Handler{
		Name:         "TM Community",
		ShortName:    "TMC",
		Description:  "TM Community solution",
		Title:        "TMC Community",
		LoadingLabel: " ",
		Icon: app.Icon{
			Default: "/web/images/loading.png",
			SVG:     "/web/images/favicon.svg",
		},
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
		Styles:          []string{"/web/styles/main.css"},
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
