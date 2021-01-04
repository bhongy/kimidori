package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// New instantiates a new Postgres connection
func New() (*pgx.Conn, error) {
	// TODO: use sslmode=verify-full
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	return newConn(dsn)
}

// NewTestDB instantiates a new Postgres connection for tests
func NewTestDB() (*pgx.Conn, error) {
	// TODO: use sslmode=verify-full (yes, for prod-test parity)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_TEST_USER"),
		os.Getenv("POSTGRES_TEST_PASSWORD"),
		os.Getenv("POSTGRES_TEST_DB"),
	)
	return newConn(dsn)
}

// newConn creates a new pgx.Conn from dsn string
func newConn(dsn string) (*pgx.Conn, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to db: %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot ping db: %v", err)
	}

	return conn, nil
}
