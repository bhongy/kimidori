package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bhongy/kimidori/authentication/internal/data"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// hardcode path relative to kimidori/authentication root for now
	// until I can figure out how to do this more intelligently
	dir := "internal/data/migrations"
	m, err := newMigrate(dir)
	if err != nil {
		log.Fatalf("Cannot create Migrate: %v\n", err)
	}

	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	err = exec(m, cmd)
	if err != nil {
		log.Fatalf("Error from executing command %q: %v\n", cmd, err)
	}

	switch cmd {
	case "up", "down", "to":
		printVersion(m)
	}
}

// newMigrate instantiates a new migreate.Migrate using postgres driver
// for migration files in `dir`
func newMigrate(dir string) (m *migrate.Migrate, err error) {
	driver, err := postgres.WithInstance(data.Db, &postgres.Config{})
	if err != nil {
		return
	}
	m, err = migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", dir),
		"postgres", driver)
	return
}

// exec runs the migration command (cmd)
func exec(m *migrate.Migrate, cmd string) (err error) {
	help := `
Usage:
	migrations current
	migrations up
	migrations down
	migrations to [version]
	`

	switch cmd {
	case "current":
		err = printVersion(m)
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "to":
		err = execTo(m)
	default:
		fmt.Println(help)
	}

	// do not consider migrate.ErrorNoChange as an error
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return
}

// execTo handles the "to" command
func execTo(m *migrate.Migrate) error {
	if len(os.Args) < 3 {
		return errors.New("No version argument provided")
	}
	arg := os.Args[2]
	v, err := strconv.ParseUint(arg, 10, 16)
	if err != nil {
		return fmt.Errorf("Cannot parse version from arg (%v): %v", arg, err)
	}
	return m.Migrate(uint(v))
}

// printVersion prints the current migration version to stdout
func printVersion(m *migrate.Migrate) error {
	v, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("Cannot read version: %v", err)
	}
	if dirty {
		fmt.Printf("version: %v (dirty)\n", v)
	} else {
		fmt.Printf("version: %v\n", v)
	}
	return nil
}
