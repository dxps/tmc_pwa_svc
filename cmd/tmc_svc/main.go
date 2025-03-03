package main

import (
	"context"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/dxps/tmc-pwa/internal/shared/config"
	"github.com/dxps/tmc-pwa/internal/svc/api"
	"github.com/dxps/tmc-pwa/internal/svc/repos"
	"github.com/dxps/tmc-pwa/internal/svc/run"
	"github.com/dxps/tmc-pwa/internal/ui/server"

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

	var cfg config.Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		slog.Error("Failed to load config.", "error", err)
		return
	}
	slog.Debug("Config loaded.")

	//////////////////////////////////////
	// API & PWA servers init & startup //
	//////////////////////////////////////

	repos, err := repos.NewRepos(
		cfg.Db.Driver,
		cfg.Db.DSN,
		cfg.Db.MaxOpenConns,
		cfg.Db.MaxIdleConns,
		cfg.Db.MaxIdleTime,
	)
	if err != nil {
		slog.Error("Failed to open database.", "error", err)
		return
	}
	defer repos.Close()

	apiSrv := api.NewApiServer(cfg.Servers.BackendPort, repos)
	apiSrv.Start()
	slog.Info("Web API Server started.", "port", apiSrv.Port)

	uiSrv := server.InitAndStartWebUiServerSide(cfg.Servers.FrontendPort, cfg.Servers.BackendPort)
	slog.Info("Web UI Server started.", "addr", uiSrv.Addr)

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

	if err := apiSrv.Shutdown(apiSrvCtx); err != nil {
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
}
