package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/dxps/tmc-pwa/internal/run"
	"github.com/dxps/tmc-pwa/internal/svc"
	"github.com/dxps/tmc-pwa/internal/svc/repos"
	"github.com/dxps/tmc-pwa/internal/ui"

	"github.com/sethvargo/go-envconfig"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				s := a.Value.Any().(*slog.Source)
				s.File = path.Base(s.File)
			}
			return a
		},
	}))
	slog.SetDefault(logger)

	slog.Info("Starting up ...")

	var cfg config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		slog.Error("Failed to load config.", "error", err)
		return
	}
	slog.Debug("Config loaded.")

	/////////////////////////////
	// db connection pool init //
	/////////////////////////////

	repos, err := repos.NewRepos(cfg.Db.Driver, cfg.Db.DSN, 10, 1, "5m")
	if err != nil {
		slog.Error("Failed to connect to database.", "error", err)
		return
	} else {
		slog.Info("Database connection established.")
	}

	/////////////////////////////////
	// http servers init & startup //
	/////////////////////////////////

	apiSrv := svc.StartApiServer(cfg.Servers.BackendPort, repos)
	slog.Info(fmt.Sprintf("Web API Server started and it's accessible at http://localhost:%d", cfg.Servers.BackendPort))

	uiSrv := ui.StartWebUiServer(cfg.Servers.FrontendPort, cfg.Servers.BackendPort)
	slog.Info(fmt.Sprintf("Web UI Server started and it's accessible at http://localhost:%d", cfg.Servers.FrontendPort))

	///////////////////////
	// graceful shutdown //
	///////////////////////

	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())
	defer shutdownCancel()

	done := run.NewOsSignalNotifier(shutdownCtx)
	<-done.Done()
	slog.Info("Shutting down ...")

	// Give outstanding requests a deadline for completion on both API and UI servers.

	apiSrvCtx, apiSrvCancel := context.WithTimeout(shutdownCtx, 3*time.Second)
	defer apiSrvCancel()

	if err := apiSrv.Stop(apiSrvCtx); err != nil {
		slog.Error("Failed to gracefully shutdown the Web API Server.", "error", err)
	} else {
		slog.Info("Web API Server gracefully shutted down.")
	}

	uiSrvCtx, uiSrvCancel := context.WithTimeout(shutdownCtx, 3*time.Second)
	defer uiSrvCancel()

	if err := uiSrv.Shutdown(uiSrvCtx); err != nil {
		slog.Error("Failed to gracefully shutdown the Web UI Server.", "error", err)
	} else {
		slog.Info("Web UI Server gracefully shutted down.")
	}

	repos.Stop()

}
