package api

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dxps/tmc-pwa/internal/svc/logic"
	"github.com/dxps/tmc-pwa/internal/svc/repos"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	server           *http.Server
	Port             int
	attributeDefMgmt *logic.AttributeDefMgmt
}

func NewApiServer(port int, repos *repos.Repos) *ApiServer {

	apiSrv := ApiServer{
		Port:             port,
		attributeDefMgmt: logic.NewAttributeDefMgmt(repos.AttributeDefRepo),
	}
	router := routerInit(&apiSrv)
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	apiSrv.server = &server

	return &apiSrv
}

// Gracefully shutdown the API Server.
func (s *ApiServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *ApiServer) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func routerInit(s *ApiServer) *chi.Mux {

	cors := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		MaxAge:         600, // Maximum value not ignored by any of the major browsers.
	})
	r := chi.NewRouter()

	// Middlewares
	r.Use(cors)
	r.Use(middleware.Logger)

	// Routes
	r.Get("/api/health", getHealthCheck)
	r.Get("/api/definitions/attributes", s.getAttributeDefs)

	return r
}
