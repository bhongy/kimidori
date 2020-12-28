package data

import (
	"database/sql"
	"fmt"
	"time"
)

// SQLWriter provides an interface to inject dependencies
// for writing data to the underlying SQL database
type SQLWriter interface {
	DB() *sql.DB
	NewUUID() string
	Now() time.Time
}

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
