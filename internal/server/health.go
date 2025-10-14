package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/alcb1310/bca-auth/internal/auth"
)

func (s *Server) CheckHealth(w http.ResponseWriter, r *http.Request) {
	_, err := auth.UserFromRequest(r)
	if err != nil {
		slog.Error("error getting user from request", "error", err)
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "unauthorized", "error": err.Error()})
		return
	}
	m := s.DB.Health()

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(m)
}
