package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

var (
	ErrDuplicateID    = errors.New("user: duplidate ID")
	ErrUsernameExists = errors.New("user: username already exists")
)

// Service provides high-level functionality
// it is agnostic to peristent, transport, or any specific I/O stack
type Service interface {
	Signup(username, password string) (User, error)
	// Login
}

// service implements business logic for the Service interface
type service struct {
	userRepo Repository
}

// NewService ...
func NewService(userRepo Repository) Service {
	return &service{userRepo}
}

// Lookup returns the value for key or ok=false if there is no mapping for key.

// Request represents a request to run a command.
// type Request struct { ...

// Encode writes the JSON encoding of req to w.
// func Encode(w io.Writer, req *Request) { ...

// Signup creates a new User
// if the username is already used, ErrNotFound is returned
func (svc *service) Signup(username, password string) (User, error) {
	var err error

	// TODO: validate username e.g. not empty, have a certain length

	// cannot signup with a username that's already used by a user
	_, err = svc.userRepo.FindByUsername(username)
	switch {
	case err == nil: // found a user
		return User{}, ErrUsernameExists
	case !errors.Is(err, ErrNotFound): // other errors
		return User{}, fmt.Errorf("signup find by username: %v", err)
	}

	p, err := NewPassword(password)
	if err != nil {
		return User{}, fmt.Errorf("hash password: %v", err)
	}

	timestamp := NewTimestamp(time.Now())
	u := User{
		ID:        ksuid.New().String(),
		Username:  username,
		Password:  p,
		CreatedAt: timestamp,
	}
	err = svc.userRepo.Create(u)

	if err != nil {
		return User{}, fmt.Errorf(
			"userRepo.Create (username=%q, password=%q): %v",
			username, password, err,
		)
	}
	return u, nil
}
