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
	"github.com/jackc/pgx/v5"
)

func (s *Server) getAllProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := s.DB.GetAllProjects()
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

func proyectValidate(p map[string]any) (types.Project, error) {
	project := types.Project{}
	var ok bool
	if project.Name, ok = p["name"].(string); !ok {
		return project, errors.New("Ingrese un nombre")
	}

	if project.GrossArea, ok = p["gross_area"].(float64); !ok {
		return project, errors.New("El area bruta debe ser un numero")
	}

	if project.NetArea, ok = p["net_area"].(float64); !ok {
		return project, errors.New("El area neta debe ser un numero")
	}

	if project.IsActive, ok = p["is_active"].(bool); !ok {
		return project, errors.New("El estado debe ser verdadero o falso")
	}

	return project, nil
}

func (s *Server) createProject(w http.ResponseWriter, r *http.Request) {
	p := map[string]any{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		slog.Error("createProject decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	project, err := proyectValidate(p)
	if err != nil {
		slog.Error("createProject validate", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	if err := s.DB.CreateProject(project); err != nil {
		if strings.Contains(err.Error(), "23505") {
			slog.Error("createProject ya existe", "err", err)
			errResponse := map[string]string{
				"error": "El proyecto ya existe",
			}
			w.WriteHeader(http.StatusConflict)
			_ = json.NewEncoder(w).Encode(errResponse)
			return
		}

		slog.Error("createProject otro error", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) updateProject(w http.ResponseWriter, r *http.Request) {
	pid := chi.URLParam(r, "id")

	projectID, err := uuid.Parse(pid)
	if err != nil {
		slog.Error("getProject parse uuid", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	if _, err = s.DB.GetProject(projectID); err != nil {
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

	p := map[string]any{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		slog.Error("createProject decode body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	project, err := proyectValidate(p)
	if err != nil {
		slog.Error("createProject validate", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}
	project.ID = projectID
	if err := s.DB.UpdateProject(project); err != nil {
		if strings.Contains(err.Error(), "23505") {
			slog.Error("createProject ya existe", "err", err)
			errResponse := map[string]string{
				"error": "El proyecto ya existe",
			}
			w.WriteHeader(http.StatusConflict)
			_ = json.NewEncoder(w).Encode(errResponse)
			return
		}

		slog.Error("createProject otro error", "err", err)
		errResponse := map[string]string{
			"error": err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(errResponse)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
