package mock_test

import (
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/repository/mock"
	"github.com/bhongy/kimidori/authentication/user"
)

var (
	repo = mock.NewUserRepository()
	now  = time.Now().Truncate(time.Millisecond)
	u    = user.User{
		ID:        "fake_user_id",
		Username:  "fake_username",
		Password:  "fake_password",
		CreatedAt: now,
	}
)

func TestUserRepository_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err := repo.Create(u)
		if err != nil {
			t.Fatal("create user: expect error to be <nil>")
		}
	})

	t.Run("failure duplicate ID", func(t *testing.T) {
		repo.Create(u)
		err := repo.Create(user.User{
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
	t.Run("found", func(t *testing.T) {
		got, err := repo.FindByUsername(u.Username)
		if err != nil {
			t.Fatal("find user by username: expect error to be <nil>")
		}
		if got != u {
			t.Errorf("Saved user mistmatch with input, got: %+v", got)
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
