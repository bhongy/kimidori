package user

// Repository provides an interface to the underlying data source
// it exposes simple CRUD operations without business logic
type Repository interface {
	Create(u *User) error

	// Create
	// GetByID
	// GetByUsername

	// FindByID
	// Store

	// FindByUsernameAndPassword
	// DoesEmailExist

	// Update
	// ChangeUsername
	// ChangePassword
}
