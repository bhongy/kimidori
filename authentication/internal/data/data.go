/*
Package `data` interfaces the underlying SQL database.

No code should interact with the database directly
without going through this package.
*/
package data

import (
	"database/sql"
	"time"
)

// SQLWriter provides an interface to inject dependencies
// for writing data to the underlying SQL database
type SQLWriter interface {
	DB() *sql.DB
	NewUUID() string
	Now() time.Time
}
