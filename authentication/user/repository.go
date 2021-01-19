package user

import "errors"

var (
	ErrNotFound = errors.New("user: not found")
)

// Repository provides an interface to the underlying data source
// it exposes simple CRUD operations without business logic
type Repository interface {
	Create(u User) error

	// FindByUsername returns User if found
	// or ErrNotFound if not found
	// otherwise returns all other error
	FindByUsername(username string) (User, error)

	// FindByID
	// FindByUsernameAndPassword
	// DoesUsernameExist

	// Update
	// ChangeUsername
	// ChangePassword
}
