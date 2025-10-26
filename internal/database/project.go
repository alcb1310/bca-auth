package database

import (
	"log/slog"

	"github.com/alcb1310/bca-auth/internal/types"
	"github.com/google/uuid"
)

func (s *service) GetAllProjects() (projects []types.Project, err error) {
	sql := "SELECT id, name, is_active, gross_area, net_area FROM project order by name"

	rows, err := s.db.Query(sql)
	if err != nil {
		slog.Error("Error executing query", "err", err)
		return
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var p types.Project
		if err = rows.Scan(&p.ID, &p.Name, &p.IsActive, &p.GrossArea, &p.NetArea); err != nil {
			slog.Error("Error scanning row", "err", err)
			return
		}
		projects = append(projects, p)
	}

	return
}

func (s *service) GetProject(id uuid.UUID) (project types.Project, err error) {
	sql := "SELECT id, name, is_active, gross_area, net_area FROM project WHERE id = $1"
	err = s.db.QueryRow(sql, id).Scan(&project.ID, &project.Name, &project.IsActive, &project.GrossArea, &project.NetArea)
	return
}

func (s *service) CreateProject(p types.Project) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("Error creating transaction", "err", err)
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	sql := "INSERT INTO project (name, is_active, gross_area, net_area) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(sql, p.Name, p.IsActive, p.GrossArea, p.NetArea)

	if err != nil {
		slog.Error("Error executing query", "err", err)
	}

	return
}

func (s *service) UpdateProject(p types.Project) (err error) {
	tx, err := s.db.Begin()
	if err != nil {
		slog.Error("Error creating transaction", "err", err)
		return
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	sql := "UPDATE project SET name = $1, is_active = $2, gross_area = $3, net_area = $4 WHERE id = $5"
	_, err = tx.Exec(sql, p.Name, p.IsActive, p.GrossArea, p.NetArea, p.ID)

	return
}
