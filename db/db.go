package db

import (
	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const Memory = ":memory:"

func NewDB(dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "db: connect to db")
	}

	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "db: setup driver for migration")
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations/", "", driver)
	if err != nil {
		return nil, errors.Wrap(err, "db: setup migration")
	}

	if err := m.Up(); err != nil {
		return nil, errors.Wrap(err, "db: run migration")
	}

	return db, err
}
