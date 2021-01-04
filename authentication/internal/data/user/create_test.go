package user_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/bhongy/kimidori/authentication/internal/data/db"
	"github.com/bhongy/kimidori/authentication/internal/data/user"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

func newTestDB() *pgx.Conn {
	db, err := db.NewTestDB()
	if err != nil {
		log.Fatalf("newTestDB: %v\n", err)
	}
	// ensure database is clean before each tests
	clean(db)
	return db
}

func clean(db *pgx.Conn) {
	// cannot use TRUNCATE TABLE due to foreign key constraint
	q := "DELETE FROM users;"
	_, err := db.Exec(context.Background(), q)
	if err != nil {
		log.Fatalf("Cannot clean db: %v\n", err)
	}
}

// findUserByID queries the database for the user for a given ID
func findUserByID(db *pgx.Conn, userID int) (u user.User, err error) {
	q := "SELECT uuid, username, created_at FROM users WHERE id = $1"
	err = db.
		QueryRow(context.Background(), q, userID).
		Scan(&u.UUID, &u.Username, &u.CreatedAt)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func NewTestUser() user.User {
	password, err := user.NewPassword("test_password")
	if err != nil {
		log.Fatalf("Error creating password: %v", err)
	}

	return user.User{
		UUID:      uuid.New(),
		Username:  "test_username",
		Password:  password,
		CreatedAt: time.Now(),
	}
}

func TestUser_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db := newTestDB()
		repo := user.NewRepository(db)
		defer repo.Close()

		u := NewTestUser()
		newUser := user.User{
			UUID:      u.UUID,
			Username:  u.Username,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
		}

		err := repo.Create(&newUser)
		if err != nil {
			t.Errorf("Expect no error but got: %v", err)
		}

		if newUser.ID == 0 {
			t.Error("Expect user id to be set")
		}

		uu, err := findUserByID(db, newUser.ID)
		if err != nil {
			t.Errorf("Cannot query user (ID=%d): %v", newUser.ID, err)
		}

		if want, got := u.UUID, uu.UUID; want != got {
			t.Errorf("Expect UUID to be %q but got %q", want, got)
		}
		if want, got := u.Username, uu.Username; want != got {
			t.Errorf("Expect Username to be %q but got %q", want, got)
		}
		// if want, got := u.CreatedAt, uu.CreatedAt; want != got {
		// 	t.Errorf("Expect CreatedAt to be %q but got %q", want, got)
		// }
	}) // t.Run("success", ...)

	t.Run("failed", func(t *testing.T) {
		db := newTestDB()
		repo := user.NewRepository(db)
		defer repo.Close()

		u := NewTestUser()
		newUser := user.User{
			UUID:      u.UUID,
			Username:  u.Username,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
		}

		repo.Create(&newUser)
		newUser.ID = 0
		// create the same user should produce an error
		err := repo.Create(&newUser)
		if err == nil {
			t.Error(errors.New("Expect error when creating the same user twice"))
		}

		if got := newUser.ID; got != 0 {
			t.Errorf("Expect User.ID to be 0 but got: %v", got)
		}
	}) // t.Run("failed", ...)
}
