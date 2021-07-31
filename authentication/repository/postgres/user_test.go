package postgres_test

import (
	"errors"
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/repository/postgres"
	"github.com/bhongy/kimidori/authentication/repository/postgres/testdb"
	"github.com/bhongy/kimidori/authentication/user"
	"github.com/google/go-cmp/cmp"
)

var (
	// call `.Truncate` since the time is stored in db with a lower precision
	// otherwise the assertion will result in a mismatch (in microseconds).
	now = user.NewTimestamp(time.Now())
	u   = user.User{
		ID:        "test_id",
		Username:  "test_username",
		Password:  "test_password",
		CreatedAt: now,
	}
)

// setup creates a new instance of user.Repository
// and ensure to "reset" db state after the current test scope (t) finishes
func setup(t *testing.T) user.Repository {
	t.Cleanup(func() {
		err := testdb.Reset(conn)
		if err != nil {
			t.Fatal("cleanup: testdb reset:", err)
		}
	})
	return postgres.NewUserRepository(conn)
}

func testCreateFirstUserSuccess(t *testing.T, repo user.Repository) {
	t.Helper()
	err := repo.Create(u)
	if err != nil {
		// no point to perform other tests if this fails
		t.Fatal("create user:", err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := setup(t)
		testCreateFirstUserSuccess(t, repo)
	})

	t.Run("failure duplicate ID", func(t *testing.T) {
		repo := setup(t)
		testCreateFirstUserSuccess(t, repo)
		err := repo.Create(user.User{
			ID:        u.ID,
			Username:  "doesntmatter",
			Password:  "doesntmatter",
			CreatedAt: now,
		})
		if err == nil {
			t.Error("create user with duplicate ID: expect error but got <nil>")
		}
	})

	t.Run("failure duplicate username", func(t *testing.T) {
		repo := setup(t)
		testCreateFirstUserSuccess(t, repo)
		err := repo.Create(user.User{
			ID:        "fake_id_2",
			Username:  u.Username,
			Password:  "fake_password_2",
			CreatedAt: user.NewTimestamp(time.Now()),
		})
		if err == nil {
			t.Error("create user with duplicate username: expect error but got <nil>")
		}
	})
}

func TestUserRepository_FindByUsername(t *testing.T) {
	repo := setup(t)
	repo.Create(u) // seed with one user

	t.Run("found", func(t *testing.T) {
		got, err := repo.FindByUsername(u.Username)
		if err != nil {
			t.Error("find user:", err)
		}
		if diff := cmp.Diff(u, got); diff != "" {
			t.Errorf("found user mistmatch (-want +got):\n%s", diff)
		}
	})

	t.Run("not found", func(t *testing.T) {
		got, err := repo.FindByUsername("this-user-should-not-exist")
		if !errors.Is(err, user.ErrNotFound) {
			t.Error("expect error to be `user.ErrNotFound` but got:", err)
		}
		if got != (user.User{}) {
			t.Errorf("expect empty User but got: %+v", got)
		}
	})
}
