package user_test

import (
	"errors"
	"testing"

	"github.com/bhongy/kimidori/authentication/repository/mock"
	"github.com/bhongy/kimidori/authentication/user"
)

const (
	username = "unicorn"
	password = "rainbow-sprinkles"
)

func setup() user.Service {
	repo := mock.NewUserRepository()
	return user.NewService(repo)
}

func TestService_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := setup()
		u, err := svc.Signup(username, password)

		if err != nil {
			t.Fatalf("Service.Signup(%q, %q): %v", username, password, err)
		}

		if u.Password.String() == password {
			t.Error("password is in plain text")
		}
	})

	t.Run("success multiple", func(t *testing.T) {
		svc := setup()
		// test that we can create multiple users with the same password
		u1, _ := svc.Signup("user1", password)
		u2, err := svc.Signup("user2", password)

		if err != nil {
			t.Errorf("create second user: %v", err)
		}

		// different users using the same password should have different hashes
		if p1, p2 := u1.Password.String(), u2.Password.String(); p1 == p2 {
			t.Errorf("same password hashes twice to the same value: %q", p1)
		}
	})

	t.Run("failure duplicate username", func(t *testing.T) {
		svc := setup()
		// create the first user
		svc.Signup(username, "password_1")
		// create a new user with the same username
		_, err := svc.Signup(username, "password_2")

		if !errors.Is(err, user.ErrUsernameExists) {
			t.Error("expect error to be user.ErrUsernameExists but got:", err)
		}
	})
}
