package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/alcb1310/bca-auth/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
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

func proveedoresValidate(p map[string]any) (types.Proveedor, error) {
	var proveedor types.Proveedor
	var ok bool

	if proveedor.Name, ok = p["name"].(string); !ok || proveedor.Name == "" {
		return proveedor, errors.New("Ingrese un nombre")
	}

	if proveedor.SupplierID, ok = p["supplier_id"].(string); !ok || proveedor.SupplierID == "" {
		return proveedor, errors.New("Ingrese un RUC")
	}

	var val string
	if val, ok = p["contact_name"].(string); !ok || val == "" {
		proveedor.ContactName = nil
	} else {
		proveedor.ContactName = &val
	}

	var val2 string
	if val2, ok = p["contact_email"].(string); !ok || val2 == "" {
		proveedor.ContactEmail = nil
	} else {
		proveedor.ContactEmail = &val2
	}

	var val3 string
	if val3, ok = p["contact_phone"].(string); !ok || val3 == "" {
		proveedor.ContactPhone = nil
	} else {
		proveedor.ContactPhone = &val3
	}

	return proveedor, nil
}

func (s *Server) createProveedores(w http.ResponseWriter, r *http.Request) {
	p := map[string]any{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		slog.Error("createProveedores decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	proveedor, err := proveedoresValidate(p)
	if err != nil {
		slog.Error("createProveedores validate", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := s.DB.CreateProveedor(proveedor); err != nil {
		if strings.Contains(err.Error(), "23505") {
			slog.Error("el proveedor ya existe", "err", err)
			errResponse := map[string]string{
				"error": "El proveedor ya existe",
			}
			w.WriteHeader(http.StatusConflict)
			_ = json.NewEncoder(w).Encode(errResponse)
			return
		}

		slog.Error("createProveedores", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) updateProveedores(w http.ResponseWriter, r *http.Request) {
	pid := chi.URLParam(r, "id")

	proveedorID, err := uuid.Parse(pid)
	if err != nil {
		slog.Error("updateProveedores parse uuid", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	p := map[string]any{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		slog.Error("updateProveedores decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	proveedor, err := proveedoresValidate(p)
	if err != nil {
		slog.Error("updateProveedores validate", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
	proveedor.ID = proveedorID

	if err := s.DB.UpdateProveedor(proveedor); err != nil {
		slog.Error("updateProveedores", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
