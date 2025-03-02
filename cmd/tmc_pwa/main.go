package main

import (
	"context"
	"log/slog"
	"os"
	"path"

	"github.com/dxps/tmc-pwa/internal/config"
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

	///////////////////////////////
	// PWA server init & startup //
	///////////////////////////////

	server.InitAndStartWebUiClientSide(cfg.Servers.FrontendPort, cfg.Servers.BackendPort)
}
