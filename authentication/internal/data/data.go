/*
Package `data` interfaces the underlying SQL database.

No code should interact with the database directly
without going through this package.
*/
package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// make "postgres" driver available for `sql.Open`
	_ "github.com/lib/pq"
)

// SQLWriter provides an interface to inject dependencies
// for writing data to the underlying SQL database
type SQLWriter interface {
	DB() *sql.DB
	NewUUID() string
	Now() time.Time
}

// Db is the singleton SQL db connection
var Db *sql.DB

func init() {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v (%v)\n", err, dsn)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping the database: %v\n", err)
	}
}
