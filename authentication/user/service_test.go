package user_test

import (
	"errors"
	"testing"

	"github.com/bhongy/kimidori/authentication/repository/mock"
	"github.com/bhongy/kimidori/authentication/user"
)

func setup() user.Service {
	repo := mock.NewUserRepository()
	return user.NewService(repo)
}

func TestService_Create(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		svc := setup()

		username := "unicorn"
		password := "rainbow-sprinkles"

		u, err := svc.Signup(username, password)
		if err != nil {
			t.Errorf("expect no error: %v", err)
		}

		if u.Password.String() == password {
			t.Fatal("expect the password to be hashed")
		}
	})

	t.Run("failure duplicate username", func(t *testing.T) {
		svc := setup()
		username := "unicorn"

		// create the first user
		svc.Signup(username, "password_1")
		// create a new user with the same username
		_, err := svc.Signup(username, "password_2")
		if !errors.Is(err, user.ErrUsernameExists) {
			t.Errorf("expect error with user.ErrUsernameExists but got: %v", err)
		}
	})

	t.Run("success multiple", func(t *testing.T) {
		svc := setup()
		// test that we can create multiple users with the same password
		password := "rainbow-sprinkles"

		u1, _ := svc.Signup("user1", password)
		u2, err := svc.Signup("user2", password)
		if err != nil {
			t.Errorf("create second user: %v", err)
		}

		if h1, h2 := u1.Password.String(), u2.Password.String(); h1 == h2 {
			t.Errorf("same password hashes twice to the same value: %v", h1)
		}
	})
}
