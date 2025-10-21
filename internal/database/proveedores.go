package database

import (
	"log/slog"

	"github.com/alcb1310/bca-auth/internal/types"
)

func (s *service) GetAllProveedores() ([]types.Proveedor, error) {
	var proveedores = []types.Proveedor{}
	sql := "SELECT id,  supplier_id, name, contact_name, contact_email, contact_phone FROM supplier order by name"

	rows, err := s.db.Query(sql)
	if err != nil {
		slog.Error("Error executing query", "err", err)
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var p types.Proveedor
		if err = rows.Scan(&p.ID, &p.SupplierID, &p.Name, &p.ContactName, &p.ContactEmail, &p.ContactPhone); err != nil {
			slog.Error("Error scanning row", "err", err)
			return nil, err
		}
		proveedores = append(proveedores, p)
	}

	return proveedores, nil
}
