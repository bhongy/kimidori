package mock

import (
	"fmt"

	"github.com/bhongy/kimidori/authentication/user"
)

// userRepository implements user.Repository for tests
type userRepository struct {
	byID       map[string]user.User
	byUsername map[string]user.User
}

func NewUserRepository() user.Repository {
	return &userRepository{
		byID:       make(map[string]user.User),
		byUsername: make(map[string]user.User),
	}
}

func (repo *userRepository) Create(u user.User) error {
	if _, ok := repo.byID[u.ID]; ok {
		return fmt.Errorf("ID (%s) already exists", u.ID)
	}
	repo.byID[u.ID] = u
	repo.byUsername[u.Username] = u
	return nil
}

func (repo *userRepository) FindByUsername(username string) (user.User, error) {
	u, ok := repo.byUsername[username]
	if !ok {
		return u, user.ErrNotFound
	}
	return u, nil
}
