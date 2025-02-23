package svc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dxps/tmc-pwa/internal/svc/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	server *http.Server
	Port   int
}

// Gracefully shutdown the API Server.
func (s *ApiServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func StartApiServer(port int) *ApiServer {

	router := routerInit()
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	apiSrv := ApiServer{
		server: &server,
		Port:   port,
	}

	go func() {
		if err := apiSrv.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	return &apiSrv
}

func routerInit() *chi.Mux {

	cors := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         600, // Maximum value not ignored by any of the major browsers.
	})
	r := chi.NewRouter()

	// Middlewares
	r.Use(cors)
	r.Use(middleware.Logger)

	// Routes
	r.Get("/health", handlers.GetHealthCheck)

	return r
}
