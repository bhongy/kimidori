// db_test checks if we can establish the connection to the database(s)
package db_test

import (
	"context"
	"testing"

	"github.com/bhongy/kimidori/authentication/internal/data/db"
)

func TestNew(t *testing.T) {
	db, err := db.New()
	defer db.Close(context.Background())
	if err != nil {
		t.Fatalf("Cannot connect to db: %v\n", err)
	}
}

func TestNewTestDB(t *testing.T) {
	db, err := db.NewTestDB()
	defer db.Close(context.Background())
	if err != nil {
		t.Fatalf("Cannot connect to test db: %v\n", err)
	}
}
