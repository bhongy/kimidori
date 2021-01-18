package user

import "time"

// User models a user account in an authentication context.
type User struct {
	ID        string
	Username  string
	Password  password // a hashed password as stored in the database
	CreatedAt time.Time
}
