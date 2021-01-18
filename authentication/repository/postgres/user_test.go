package postgres_test

import (
	"context"
	"errors"
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

// reset deletes all rows in the related table from the database
func reset(conn *pgx.Conn) {
	// cannot use TRUNCATE TABLE due to foreign key constraint
	q := "DELETE from users"
	_, err := conn.Exec(context.Background(), q)
	if err != nil {
		log.Fatalf("cannot exec users delete query: %v", err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	conn, err := testdb.Open()
	defer conn.Close(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	repo := postgres.NewUserRepository(conn)
	// Truncate since the time stored in DB has a lower precision
	// otherwise `now` here and the value retrieves from the db later
	// won't match
	now := time.Now().Truncate(time.Millisecond)

	t.Run("success", func(t *testing.T) {
		in := user.User{
			ID:        "stub_id",
			Username:  "stub_username",
			Password:  "stub_password",
			CreatedAt: now,
		}
		err = repo.Create(in)
		defer reset(conn)
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
	})

	t.Run("failure duplicate ID", func(t *testing.T) {
		id := "stub_user_id"
		repo.Create(user.User{
			ID:        id,
			Username:  "name1",
			Password:  "pass1",
			CreatedAt: now,
		})
		err := repo.Create(user.User{
			ID:        id,
			Username:  "name2",
			Password:  "pass2",
			CreatedAt: now,
		})
		defer reset(conn)
		if err == nil {
			t.Error("expect error but got <nil>")
		}
	})

	t.Run("failure duplicate username", func(t *testing.T) {
		username := "stub_username"
		repo.Create(user.User{
			ID:        "id1",
			Username:  username,
			Password:  "pass1",
			CreatedAt: now,
		})
		defer reset(conn)
		err := repo.Create(user.User{
			ID:        "id2",
			Username:  username,
			Password:  "pass2",
			CreatedAt: now,
		})
		if err == nil {
			t.Error("expect error but got <nil>")
		}
	})
}

func TestUserRepository_FindByUsername(t *testing.T) {
	ctx := context.Background()
	conn, err := testdb.Open()
	defer conn.Close(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	repo := postgres.NewUserRepository(conn)
	// Truncate since the time stored in DB has a lower precision
	// otherwise `now` here and the value retrieves from the db later
	// won't match
	now := time.Now().Truncate(time.Millisecond)
	u := user.User{
		ID:        "stub_id",
		Username:  "stub_username",
		Password:  "stub_password",
		CreatedAt: now,
	}
	// seed with one user
	err = repo.Create(u)
	defer reset(conn)
	if err != nil {
		log.Fatalf("create user: %v\n", err)
	}

	t.Run("success", func(t *testing.T) {
		got, err := repo.FindByUsername(u.Username)
		if err != nil {
			t.Fatalf("expect no error but got: %v", err)
		}
		if diff := cmp.Diff(u, got); diff != "" {
			t.Errorf("saved user mistmatch with input (-want +got):\n%s", diff)
		}
	})

	t.Run("failure not found", func(t *testing.T) {
		got, err := repo.FindByUsername("this-user-should-not-exist")
		if !errors.Is(err, user.ErrNotFound) {
			t.Errorf("expect error to be `user.ErrNotFound` but got: %v", err)
		}
		if got != (user.User{}) {
			t.Errorf("expect empty User but got: %v", got)
		}
	})
}
