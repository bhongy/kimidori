package mock

import (
	"fmt"

	"github.com/bhongy/kimidori/authentication/user"
)

// userRepository implements user.Repository for tests.
type userRepository struct {
	// use pointer to keep both maps in sync
	byID       map[string]*user.User
	byUsername map[string]*user.User
}

// NewUserRepository creates a new user.Repository for tests using in-memory store.
func NewUserRepository() user.Repository {
	return &userRepository{
		byID:       make(map[string]*user.User),
		byUsername: make(map[string]*user.User),
	}
}

func (repo *userRepository) Create(u user.User) error {
	if _, ok := repo.byID[u.ID]; ok {
		return fmt.Errorf("ID (%s) already exists", u.ID)
	}
	p := &u
	repo.byID[u.ID] = p
	repo.byUsername[u.Username] = p
	return nil
}

func (repo *userRepository) FindByUsername(username string) (user.User, error) {
	p := repo.byUsername[username]
	if p == nil {
		return user.User{}, user.ErrNotFound
	}
	return *p, nil
}
