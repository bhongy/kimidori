package postgres

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// PostgresWriter implements SQLWriter to interface with Postgres
type PostgresWriter struct {
	db *sql.DB
}

func (w PostgresWriter) DB() *sql.DB {
	return w.db
}

func (w PostgresWriter) NewUUID() string {
	return uuid.New().String()
}

func (w PostgresWriter) Now() time.Time {
	return time.Now()
}

// NewPostgresWriter instantiates a new PostgresWriter
// using the `db` connection
func NewPostgresWriter(db *sql.DB) PostgresWriter {
	return PostgresWriter{db}
}
