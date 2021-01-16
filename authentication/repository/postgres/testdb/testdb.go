package testdb

import (
	"fmt"
	"os"

	"github.com/bhongy/kimidori/authentication/repository/postgres/db"
	"github.com/jackc/pgx/v4"
)

// Open returns a new postgres connection for the test database
func Open() (*pgx.Conn, error) {
	// TODO: use sslmode=verify-full (yes, for prod-test parity)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_TEST_USER"),
		os.Getenv("POSTGRES_TEST_PASSWORD"),
		os.Getenv("POSTGRES_TEST_DB"),
	)
	return db.NewConnection(dsn)
}
