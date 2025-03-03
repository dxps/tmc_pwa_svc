package api

// initRoutes registers the UI routes.
func (s *ApiServer) initRoutes() {

	s.router.Get("/health", s.getHealthCheck)
	s.router.Get("/api/definitions/attributes", s.getAttributeDefs)
}
