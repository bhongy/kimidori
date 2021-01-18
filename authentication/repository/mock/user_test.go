package mock_test

import (
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/repository/mock"
	"github.com/bhongy/kimidori/authentication/user"
	"github.com/google/go-cmp/cmp"
)

var (
	repo user.Repository
	now  = time.Now().Truncate(time.Millisecond)
	u    = user.User{
		ID:        "fake_user_id",
		Username:  "fake_username",
		Password:  "fake_password",
		CreatedAt: now,
	}
)

func setup(t *testing.T) {
	repo = mock.NewUserRepository()
	// always clear repo after each test scope finishes
	// so it is difficult to have an accidental shared state
	t.Cleanup(func() { repo = nil })
}

func TestUserRepository_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		setup(t)
		err := repo.Create(u)
		if err != nil {
			t.Fatal("create user: expect error to be <nil>")
		}
	})

	t.Run("failure duplicate ID", func(t *testing.T) {
		setup(t)
		// TODO: find a pattern to combine this test and the success test
		err := repo.Create(u)
		if err != nil {
			t.Fatal("seeding first user:", err)
		}
		err = repo.Create(user.User{
			ID:        u.ID,
			Username:  "doesntmatter",
			Password:  "doesntmatter",
			CreatedAt: now,
		})
		if err == nil {
			t.Error("expect error but got <nil>")
		}
	})
}

func TestUserRepository_FindByUsername(t *testing.T) {
	setup(t)
	repo.Create(u)

	t.Run("found", func(t *testing.T) {
		got, err := repo.FindByUsername(u.Username)
		if err != nil {
			t.Fatal("find user by username: expect error to be <nil>")
		}
		if diff := cmp.Diff(u, got); diff != "" {
			t.Errorf("found user mistmatch (-want +got):\n%s", diff)
		}
	})

	t.Run("not found", func(t *testing.T) {
		got, err := repo.FindByUsername("this-username-should-not-exist")
		if err == nil {
			t.Error("expect error but got <nil>")
		}
		if got != (user.User{}) {
			t.Errorf("expect empty user but got: %+v", got)
		}
	})
}
