package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := s.DB.GetAllProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(projects)
}
