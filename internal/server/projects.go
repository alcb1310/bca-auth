package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
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
