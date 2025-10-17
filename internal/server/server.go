package server

import (
	"net/http"

	"github.com/alcb1310/bca-auth/internal/auth"
	"github.com/alcb1310/bca-auth/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	Router *chi.Mux
	DB     database.Service
}

func NewServer() *Server {
	s := &Server{
		Router: chi.NewRouter(),
		DB:     database.NewService(),
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
		r.Use(auth.IsAuthenticated())

		r.Get("/health", s.CheckHealth)

		r.Route("/parametros", func(r chi.Router) {
			r.Route("/partidas", func(r chi.Router) {})
			r.Route("/categorias", func(r chi.Router) {})
			r.Route("/materiales", func(r chi.Router) {})

			r.Route("/proyectos", func(r chi.Router) {
				r.Get("/", s.getAllProjects)
			})

			r.Route("/proveeodres", func(r chi.Router) {})
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
