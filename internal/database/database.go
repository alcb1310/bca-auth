package database

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/alcb1310/bca-auth/internal/types"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type service struct {
	db *sql.DB
}

type Service interface {
	Health() (map[string]string, error)

	GetAllProjects() (projects []types.Project, err error)
	GetProject(id uuid.UUID) (project types.Project, err error)
	CreateProject(p types.Project) (err error)
	UpdateProject(p types.Project) (err error)

	GetAllProveedores() (proveedores []types.Proveedor, err error)
	CreateProveedor(p types.Proveedor) (err error)
}

func NewService() Service {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		slog.Error("connecting to db", "err", err)
		panic(err)
	}

	err = createTables(db)
	if err != nil {
		slog.Error("creating tables", "err", err)
		panic(err)
	}

	slog.Info("db up")

	s := &service{
		db: db,
	}
	return s
}

func (s *service) Health() (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		slog.Error("db down", "err", err)
		return nil, err
	}

	return map[string]string{
		"message": "It's healthy",
	}, nil
}
