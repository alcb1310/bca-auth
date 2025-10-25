package server

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/alcb1310/bca-auth/internal/types"
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

	if proveedor.Name, ok = p["name"].(string); !ok {
		return proveedor, errors.New("Ingrese un nombre")
	}
	if proveedor.Name == "" {
		return proveedor, errors.New("Ingrese un nombre")
	}

	if proveedor.SupplierID, ok = p["supplier_id"].(string); !ok {
		return proveedor, errors.New("Ingrese un RUC")
	}
	if proveedor.SupplierID == "" {
		return proveedor, errors.New("Ingrese un RUC")
	}

	if proveedor.ContactName, ok = p["contact_name"].(*string); !ok {
		proveedor.ContactName = nil
	}

	if proveedor.ContactEmail, ok = p["email"].(*string); !ok {
		proveedor.ContactEmail = nil
	}

	if proveedor.ContactPhone, ok = p["contact_phone"].(*string); !ok {
		proveedor.ContactPhone = nil
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
