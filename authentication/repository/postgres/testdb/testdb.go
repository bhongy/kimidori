package testdb

import (
	"context"
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

// Reset deletes all rows in the related table from the database
func Reset(conn *pgx.Conn) error {
	// cannot use TRUNCATE TABLE due to foreign key constraint
	q := "DELETE from users"
	_, err := conn.Exec(context.Background(), q)
	if err != nil {
		return fmt.Errorf("exec delete from users: %v", err)
	}
	return nil
}
