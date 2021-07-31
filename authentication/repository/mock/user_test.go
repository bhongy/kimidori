package mock_test

import (
	"errors"
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/repository/mock"
	"github.com/bhongy/kimidori/authentication/user"
	"github.com/google/go-cmp/cmp"
)

var (
	now = user.NewTimestamp(time.Now())
	u   = user.User{
		ID:        "fake_user_id",
		Username:  "fake_username",
		Password:  "fake_password",
		CreatedAt: now,
	}
)

func setup(t *testing.T) user.Repository {
	repo := mock.NewUserRepository()
	// always clear repo after each test scope finishes
	// so it is difficult to have an accidental shared state
	t.Cleanup(func() { repo = nil })
	return repo
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
		testCreateFirstUserSuccess(t, repo) // seed with the first user
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
}

func TestUserRepository_FindByUsername(t *testing.T) {
	repo := setup(t)
	testCreateFirstUserSuccess(t, repo)

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
