package user

import "time"

type User struct {
	ID        string
	Username  string
	Password  string // a hashed password as stored in the database
	CreatedAt time.Time
}
