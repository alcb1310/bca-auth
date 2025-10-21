package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *Server) getAllProveedores(w http.ResponseWriter, r *http.Request) {
	proveedores, err := s.DB.GetAllProveedores()
	if err != nil {
		slog.Error("getAllProveedores", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(proveedores)
}
