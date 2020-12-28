package data_test

import (
	"database/sql"
	"time"
)

type testSQLWriter struct {
	db   *sql.DB
	uuid string
	now  time.Time
}

func (w testSQLWriter) DB() *sql.DB {
	return w.db
}

func (w testSQLWriter) NewUUID() string {
	return w.uuid
}

func (w testSQLWriter) Now() time.Time {
	return w.now
}
