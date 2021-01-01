package user

import (
	"context"
	"fmt"
	"time"
)

func (r repository) Create(u *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	q := `
		INSERT INTO users (uuid, username, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id
	`
	err := r.db.
		QueryRowContext(ctx, q, u.UUID, u.Username, u.Password, u.CreatedAt).
		Scan(&u.ID)
	if err != nil {
		return fmt.Errorf("create user: %v", err)
	}
	return nil
}
