package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) CheckHealth(w http.ResponseWriter, r *http.Request) {
	m, err := s.DB.Health()
	if err != nil {
		slog.Error("health", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(m)
}
