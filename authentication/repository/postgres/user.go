package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bhongy/kimidori/authentication/user"
	"github.com/jackc/pgx/v4"
)

// userRepository implements user.Repository interface
type userRepository struct {
	conn *pgx.Conn
}

// userRepo := postgres.NewUserRepository(...)

func NewUserRepository(conn *pgx.Conn) user.Repository {
	return &userRepository{conn}
}

func (repo *userRepository) Create(u user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q := `
		INSERT INTO users (id, username, password, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := repo.conn.Exec(ctx, q, u.ID, u.Username, u.Password, u.CreatedAt)
	if err != nil {
		return fmt.Errorf("create user: %v", err)
	}
	return nil
}

func (repo *userRepository) FindByUsername(username string) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u user.User
	q := `
		SELECT id, username, password, created_at
		FROM users
		WHERE username = $1
	`
	err := repo.conn.
		QueryRow(ctx, q, username).
		Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return u, user.ErrNotFound
		}
		return u, fmt.Errorf("find by username: %v", err)
	}
	return u, nil
}
