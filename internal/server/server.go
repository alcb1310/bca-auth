package server

import (
	"net/http"

	"github.com/alcb1310/bca-auth/internal/auth"
	"github.com/alcb1310/bca-auth/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Router *chi.Mux
	DB     database.Service
	User   auth.User
}

func NewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
		DB:     database.NewService(),
		User:   auth.NewUser(),
	}

	s.Router.Use(middleware.Logger)

	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from server"))
	})

	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.AllowContentType("application/json"))
		r.Use(s.User.IsAuthenticated())

		r.Get("/health", s.CheckHealth)

		r.Route("/parametros", func(r chi.Router) {
			r.Route("/partidas", func(r chi.Router) {})
			r.Route("/categorias", func(r chi.Router) {})
			r.Route("/materiales", func(r chi.Router) {})

			r.Route("/proyectos", func(r chi.Router) {
				r.Get("/", s.getAllProjects)
				r.Post("/", s.createProject)

				r.Get("/{id}", s.getProject)
				r.Put("/{id}", s.updateProject)
			})

			r.Route("/proveedores", func(r chi.Router) {
				r.Get("/", s.getAllProveedores)
				r.Post("/", s.createProveedores)
			})
			r.Route("/rubros", func(r chi.Router) {})
		})

		r.Route("/transacciones", func(r chi.Router) {
			r.Route("/presupuestos", func(r chi.Router) {})
			r.Route("/facturas", func(r chi.Router) {})
			r.Route("/cierre", func(r chi.Router) {})
		})

		r.Route("/reportes", func(r chi.Router) {
			r.Route("/actual", func(r chi.Router) {})
			r.Route("/cuadre", func(r chi.Router) {})
			r.Route("/gastado", func(r chi.Router) {})
			r.Route("/historico", func(r chi.Router) {})
		})

		r.Route("/analisis", func(r chi.Router) {
			r.Route("/cantidades", func(r chi.Router) {})
			r.Route("/analisis", func(r chi.Router) {})
		})
	})

	return s
}
