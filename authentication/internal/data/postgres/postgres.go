package postgres

import (
	"database/sql"
	"fmt"
	"os"

	// make "postgres" driver available for `sql.Open`
	_ "github.com/lib/pq"
)

// NewDB instantiates a new sql.DB connection for Postgres
func NewDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// err = fmt.Errorf("Cannot connect to the database: %v (%v)\n", err, dsn)
		return nil, fmt.Errorf("Cannot connect to the database: %v (%v)\n", err, dsn)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Cannot ping the database: %v\n", err)
	}

	return db, nil
}
