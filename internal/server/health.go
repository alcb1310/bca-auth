package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) CheckHealth(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
