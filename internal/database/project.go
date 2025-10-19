package database

import (
	"log/slog"

	"github.com/alcb1310/bca-auth/internal/types"
)

func (s *service) GetAllProjects() ([]types.Project, error) {
	var projects = []types.Project{}
	sql := "SELECT id, name, is_active, gross_area, net_area FROM project order by name"

	rows, err := s.db.Query(sql)
	if err != nil {
		slog.Error("Error executing query", "err", err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var p types.Project
		if err = rows.Scan(&p.ID, &p.Name, &p.IsActive, &p.GrossArea, &p.NetArea); err != nil {
			slog.Error("Error scanning row", "err", err)
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
