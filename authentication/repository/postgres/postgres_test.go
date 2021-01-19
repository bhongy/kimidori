package postgres_test

import (
	"context"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/bhongy/kimidori/authentication/repository/postgres/testdb"
	"github.com/jackc/pgx/v4"
)

var conn *pgx.Conn

// TestMain sets up the test environment for _all_ tests in the package
func TestMain(m *testing.M) {
	err := runMigrations()
	if err != nil {
		log.Fatalf("run migrations: %v\n", err)
	}

	// do not do `conn, err := ...` otherwise it'll create a new variable
	// instead of setting the package global `conn`
	conn, err = testdb.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// no need to call `os.Exit(m.Run())`
	// https://github.com/golang/go/issues/34129
	//
	// `os.Exit` bypasses the defered cleanups
	m.Run()
}

// runMigrations migrates the test db to latest
func runMigrations() error {
	cmd := exec.Command("tern", "migrate", "-c", "tern-testdb.conf")
	// relative path to repository/postgres/migrations package
	cmd.Dir = "./migrations"
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
