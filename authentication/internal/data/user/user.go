package user

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Close()
	Create(user *User) error
}

type User struct {
	ID       int
	UUID     uuid.UUID
	Username string
	// Password is the hashed password as stored in the database
	Password  []byte
	CreatedAt time.Time
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Close() {
	r.db.Close()
}
