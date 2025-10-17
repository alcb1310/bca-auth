package database

import (
	"database/sql"
	"log/slog"
	"os"
	"strings"
)

func createTables(d *sql.DB) error {
	file, err := os.ReadFile("./internal/database/queries/tables.sql")
	if err != nil {
		slog.Error("Error reading file", "err", err)
		return err
	}

	tx, err := d.Begin()
	if err != nil {
		slog.Error("Error creating transaction", "err", err)
		return err
	}

	requests := strings.SplitSeq(string(file), ";")
	for request := range requests {
		if _, err := tx.Exec(request); err != nil {
			slog.Error("Error executing query", "err", err)
			_ = tx.Rollback()
			return err
		}
	}
	_ = tx.Commit()

	return nil
}
