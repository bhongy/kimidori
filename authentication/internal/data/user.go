package data

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Username  string
	Password  string
	CreatedAt time.Time
}

// CreateUser creates a new user in the database
func CreateUser(w SQLWriter, username, password string) (u User, err error) {
	// TODO: encrypt password
	u = User{
		UUID:      w.NewUUID(),
		Username:  username,
		Password:  password,
		CreatedAt: w.Now(),
	}
	q := `
		INSERT INTO users (uuid, username, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id
	`
	err = w.DB().
		QueryRow(q, u.UUID, u.Username, u.Password, u.CreatedAt).
		Scan(&u.ID)
	if err != nil {
		u = User{}
		err = fmt.Errorf("create user: %v", err)
	}
	return
}

// Create creates a new user in the database
//
// it is an alternative (experimental) API to `CreateUser`
//
// this API doesn't require passing a function to produce uuid, timestamp
// but let that be the caller's concern
func (u *User) Create(db *sql.DB, plaintextPassword string) error {
	// TODO: encrypt password
	u.Password = plaintextPassword
	q := `
		INSERT INTO users (uuid, username, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id
	`
	err := db.
		QueryRow(q, u.UUID, u.Username, plaintextPassword, u.CreatedAt).
		Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("create user: %v", err)
	}
	return nil
}
