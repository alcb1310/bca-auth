package database

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type service struct {
	db *sql.DB
}

type Service interface {
	Health() (map[string]string, error)
}

func NewService() Service {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		slog.Error("connecting to db", "err", err)
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
