package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
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
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) Repository {
	return &repository{db}
}

func (r *repository) Close() {
	r.db.Close(context.Background())
}
