package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDatabase creates a new database connection pool.
func NewDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return db, nil
}
