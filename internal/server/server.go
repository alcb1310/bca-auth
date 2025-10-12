package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
}

func NewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
	}

	s.Router.Use(middleware.Logger)
	s.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Type", "Content-Disposition"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from server"))
	})

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))

		r.Get("/health", s.CheckHealth)
	})

	return s
}
