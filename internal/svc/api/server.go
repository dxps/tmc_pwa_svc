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
	router           *chi.Mux
	server           *http.Server
	Port             int
	attributeDefMgmt *logic.AttributeDefMgmt
}

func NewApiServer(port int, repos *repos.Repos) *ApiServer {

	apiSrv := ApiServer{
		Port:             port,
		attributeDefMgmt: logic.NewAttributeDefMgmt(repos.AttributeDefRepo),
	}
	apiSrv.initRouter()
	apiSrv.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: apiSrv.router,
	}

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

func (s *ApiServer) initRouter() {

	cors := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},                                       //
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, //
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"}, //
		MaxAge:         600,                                                 // Maximum value not ignored by any of the major browsers.
	})
	s.router = chi.NewRouter()

	// Middlewares setup.
	s.router.Use(cors)
	s.router.Use(middleware.Logger)

	s.initRoutes()
}
