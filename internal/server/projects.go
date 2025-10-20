package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *Server) getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := s.DB.GetAllProjects()
	slog.Info("getAllProjects", "user from context", r.Context().Value("user"))
	slog.Info("getAllProjects", "user from Server", s.User)
	if err != nil {
		slog.Error("getAllProjects", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}

func (s *Server) getProject(w http.ResponseWriter, r *http.Request) {

	pid := chi.URLParam(r, "id")

	projectID, err := uuid.Parse(pid)
	if err != nil {
		slog.Error("getProject parse uuid", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	project, err := s.DB.GetProject(projectID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		slog.Error("getProject", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(project)

}
