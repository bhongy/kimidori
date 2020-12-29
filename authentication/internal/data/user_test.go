package data_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bhongy/kimidori/authentication/internal/data"
)

// createSqlMock initializes a new db and sqlmock instances
func createSqlMock(t *testing.T) (db *sql.DB, mock sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Fail creating SQL mock: %v\n", err)
	}
	return
}

func TestCreateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock := createSqlMock(t)
		defer db.Close()

		username := "test_username"
		password := "test_password"
		w := testSQLWriter{
			db:   db,
			uuid: "test_uuid_17fc",
			now:  time.Now(),
		}

		id := 42
		rows := sqlmock.
			NewRows([]string{"id"}).
			AddRow(id)

		mock.
			ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING id").
			WithArgs(w.uuid, username, password, w.now).
			WillReturnRows(rows)

		got, err := data.CreateUser(w, username, password)
		if err != nil {
			t.Error(err)
		}

		want := data.User{
			ID:        id,
			UUID:      w.uuid,
			Username:  username,
			Password:  password,
			CreatedAt: w.now,
		}
		if got != want {
			t.Errorf("\ngot: %v\nwant: %v", got, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}) // t.Run("success", ...)

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		db, mock := createSqlMock(t)
		defer db.Close()

		username := "test_username"
		password := "test_password"
		w := testSQLWriter{
			db:   db,
			uuid: "test_uuid_17fc",
			now:  time.Now(),
		}

		mock.
			ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING id").
			WithArgs(w.uuid, username, password, w.now).
			WillReturnError(errors.New("Stub error from executing the query"))

		got, err := data.CreateUser(w, username, password)
		if err == nil {
			t.Error(errors.New("Expect error to be returned but got `nil`"))
		}

		if got != (data.User{}) {
			t.Errorf("Expect empty User but got: %v", got)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}) // t.Run("failed", ...)
}

func TestUser_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		t.Parallel()

		db, mock := createSqlMock(t)
		defer db.Close()

		u := data.User{
			UUID:      "test_uuid_17fc",
			Username:  "test_username",
			CreatedAt: time.Now(),
		}
		password := "test_password"

		id := 42
		rows := sqlmock.
			NewRows([]string{"id"}).
			AddRow(id)

		mock.
			ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING id").
			WithArgs(u.UUID, u.Username, password, u.CreatedAt).
			WillReturnRows(rows)

		if err := u.Create(db, password); err != nil {
			t.Error(err)
		}

		want := data.User{
			ID:        id,
			UUID:      u.UUID,
			Username:  u.Username,
			Password:  password,
			CreatedAt: u.CreatedAt,
		}
		if u != want {
			t.Errorf("\ngot: %v\nwant: %v", u, want)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	})

	t.Run("failed", func(t *testing.T) {
		t.Parallel()

		db, mock := createSqlMock(t)
		defer db.Close()

		u := data.User{
			UUID:      "test_uuid_17fc",
			Username:  "test_username",
			CreatedAt: time.Now(),
		}
		password := "test_password"

		mock.
			ExpectQuery("^INSERT INTO users (.+) VALUES (.+) RETURNING id").
			WithArgs(u.UUID, u.Username, password, u.CreatedAt).
			WillReturnError(errors.New("Stub error from executing the query"))

		if err := u.Create(db, password); err == nil {
			t.Error(errors.New("Expect error to be returned but got `nil`"))
		}

		if got := u.ID; got != 0 {
			t.Errorf("Expect User.ID to be 0 but got: %v", got)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unfulfilled expectations: %s", err)
		}
	}) // t.Run("failed", ...)
}
