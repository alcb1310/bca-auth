package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) CheckHealth(w http.ResponseWriter, r *http.Request) {
	m := s.DB.Health()

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(m)
}
