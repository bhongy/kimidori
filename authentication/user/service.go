package user

// Service provides high-level functionality
// it is agnostic to peristent, transport, or any specific I/O stack
type Service interface {
	// Signup
	// Login
}

// service implements business logic for the Service interface
type service struct {
}

func NewService() Service {
	return &service{}
}

// Signup
// - cannot sign up if the ID already exists
//   - retry with a different generated ID
//   - or just ensure ID will always be unique <-- this
// - cannot sign up if the user is already used
//
// ...
//
// - must provide ID, username, password
// - password is not stored in plain text
