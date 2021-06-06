package db

import (
	"context"
	"embed"

	"github.com/cockroachdb/errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const Memory = ":memory:"

//go:embed migrations/*.sql
var fs embed.FS

func NewDB(ctx context.Context, dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "sqlite3", dataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "db: connect to db")
	}

	srcDriver, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, errors.Wrap(err, "db: setup source driver for migration")
	}

	dbDriver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "db: setup database driver for migration")
	}

	m, err := migrate.NewWithInstance("iofs", srcDriver, "", dbDriver)
	if err != nil {
		return nil, errors.Wrap(err, "db: setup migration")
	}

	if err := m.Up(); err != nil {
		return db, errors.Wrap(err, "db: run migration")
	}

	return db, err
}
