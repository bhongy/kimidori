package user_test

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bhongy/kimidori/authentication/internal/data/user"
	"github.com/google/uuid"
)

var u = user.User{
	ID:        42,
	UUID:      uuid.New(),
	Username:  "test_username",
	Password:  "test_password",
	CreatedAt: time.Now(),
}

var ErrCreate = errors.New("(stub) cannot create user")

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Error creating SQL mock: %v\n", err)
	}
	return db, mock
}

func TestUser_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock := NewMock()
		repo := user.NewRepository(db)
		defer repo.Close()

		query := "INSERT INTO users (.+) VALUES (.+) RETURNING id"
		rows := sqlmock.NewRows([]string{"id"}).AddRow(u.ID)
		mock.
			ExpectQuery(query).
			WithArgs(u.UUID, u.Username, u.Password, u.CreatedAt).
			WillReturnRows(rows)

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

		want := user.User{
			ID:        u.ID,
			UUID:      u.UUID,
			Username:  u.Username,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
		}
		if newUser != want {
			t.Errorf("\ngot: %v\nwant: %v", newUser, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}) // t.Run("success", ...)

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		db, mock := NewMock()
		repo := user.NewRepository(db)
		defer repo.Close()

		query := "INSERT INTO users (.+) VALUES (.+) RETURNING id"
		mock.
			ExpectQuery(query).
			WithArgs(u.UUID, u.Username, u.Password, u.CreatedAt).
			WillReturnError(ErrCreate)

		newUser := user.User{
			UUID:      u.UUID,
			Username:  u.Username,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
		}

		err := repo.Create(&newUser)
		if err == nil {
			t.Error(errors.New("Expect error to be returned but got `nil`"))
		}

		if got := newUser.ID; got != 0 {
			t.Errorf("Expect User.ID to be 0 but got: %v", got)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}) // t.Run("failed", ...)
}
