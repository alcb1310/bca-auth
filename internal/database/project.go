package database

import "github.com/alcb1310/bca-auth/internal/types"

func (s *service) GetAllProjects() ([]types.Project, error) {
	sql := "SELECT id, name, is_active, gross_area, net_area FROM projects order by name"

	rows, err := s.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var projects []types.Project
	for rows.Next() {
		var p types.Project
		err = rows.Scan(&p.ID, &p.Name, &p.IsActive, &p.GrossArea, &p.NetArea)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}
