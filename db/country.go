package db

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
)

type CountryTable struct {
	db *sqlx.DB
}

func NewCountryTable(db *sqlx.DB) *CountryTable {
	return &CountryTable{db: db}
}

func (c *CountryTable) UpsertCountry(ctx context.Context, alpha2Code, englishShortName string) error {
	_, err := c.db.ExecContext(ctx, `
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

func (c *CountryTable) getCountryEnglishShortName(ctx context.Context, alpha2Code string) (string, error) {
	var englishShortName string

	err := c.db.QueryRowxContext(ctx, `
SELECT
    english_sort_name
FROM
    countries
WHERE
    alpha_2_code = ?
;
    `, alpha2Code).Scan(&englishShortName)

	if err != nil {
		return "", errors.Wrap(err, "db: get english_sort_name")
	}

	return englishShortName, nil
}
