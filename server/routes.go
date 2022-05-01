package server

import (
	"github.com/go-chi/chi/v5"
)

func (srv *Server) InjectRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get(`/health`, srv.healthCheck)

	router.Route("/api", func(api chi.Router) {
		api.Get("/welcome", srv.greet)
		api.Post("/upload_csv", srv.uploadCSV)
		//api.Post("/upload", srv.upload)
	})
	return router
}
