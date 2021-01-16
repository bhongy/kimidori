package mock

import (
	"fmt"

	"github.com/bhongy/kimidori/authentication/user"
)

// userRepository implements user.Repository for tests
type userRepository struct {
	// store maps User.ID to User
	store map[string]user.User
}

func NewUserRepository() user.Repository {
	store := make(map[string]user.User)
	return &userRepository{store}
}

func (repo *userRepository) Create(u user.User) error {
	if _, ok := repo.store[u.ID]; ok {
		return fmt.Errorf("ID (%s) already exists", u.ID)
	}
	repo.store[u.ID] = u
	return nil
}
