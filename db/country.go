package db

import (
	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

type CountryTable struct {
	db *sqlx.DB
}

func NewCountryTable(db *sqlx.DB) *CountryTable {
	return &CountryTable{db: db}
}

func (c *CountryTable) UpsertCountry(alpha2Code, englishShortName string) error {
	_, err := c.db.Exec(`
INSERT INTO countries (alpha_2_code, english_sort_name)
    VALUES (?, ?)
ON CONFLICT (alpha_2_code)
    DO UPDATE SET
        english_sort_name = excluded.english_sort_name;
    `, alpha2Code, englishShortName)

	if err != nil {
		return errors.Wrap(err, "db: upsert country")
	}

	return nil
}
