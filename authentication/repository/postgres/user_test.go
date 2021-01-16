package postgres_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/repository/postgres"
	"github.com/bhongy/kimidori/authentication/repository/postgres/testdb"
	"github.com/bhongy/kimidori/authentication/user"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
)

func TestMain(m *testing.M) {
	err := runMigrations()
	if err != nil {
		log.Fatalf("run migrations: %v\n", err)
	}
	os.Exit(m.Run())
}

func runMigrations() error {
	cmd := exec.Command("tern", "migrate", "-c", "tern-testdb.conf")
	// relative path to repository/postgres/migrations package
	cmd.Dir = "./migrations"
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// deleteUsers deletes all rows in the users table from the database
func deleteUsers(conn *pgx.Conn) error {
	q := "DELETE from users"
	_, err := conn.Exec(context.Background(), q)
	if err != nil {
		return fmt.Errorf("cannot exec users delete query: %v", err)
	}
	return nil
}

// TODO: test inserting to the ID already exists is an error
func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	conn, err := testdb.Open()
	defer conn.Close(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	err = deleteUsers(conn)
	if err != nil {
		log.Fatalf("deleteUsers: %v\n", err)
	}

	repo := postgres.NewUserRepository(conn)
	// Truncate since the time stored in DB has a lower precision
	// otherwise `now` here and the value retrieves from the db later
	// won't match
	now := time.Now().Truncate(time.Microsecond)
	in := user.User{
		ID:        "stub_id",
		Username:  "stub_username",
		Password:  "stub_password",
		CreatedAt: now,
	}
	err = repo.Create(in)
	if err != nil {
		t.Fatalf("repo.Create: %v", err)
	}

	var out user.User
	err = conn.
		QueryRow(ctx, `
			SELECT id, username, password, created_at
			FROM users
		`).
		Scan(&out.ID, &out.Username, &out.Password, &out.CreatedAt)
	if err != nil {
		t.Fatalf("querying created user: %v", err)
	}

	if diff := cmp.Diff(in, out); diff != "" {
		t.Errorf("Saved user mistmatch with input (-in +out):\n%s", diff)
	}
}
