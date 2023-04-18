package storage

import (
	"path/filepath"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Sqlite3Migrate(migrations, storageFile string) (err error) {
	var mabs string
	mabs, err = filepath.Abs(migrations)
	if err != nil {
		return
	}

	var dbabs string
	dbabs, err = filepath.Abs(storageFile)
	if err != nil {
		return
	}

	var sURL, dURL string
	sURL = "file://" + mabs
	dURL = "sqlite3://" + dbabs

	var m *migrate.Migrate
	m, err = migrate.New(sURL, dURL)
	if err != nil {
		return
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		err = nil
	}
	return
}
