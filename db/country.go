package db

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
)

type CountryTable struct {
	db *sqlx.DB
}

func NewCountryTable(db *sqlx.DB) *CountryTable {
	return &CountryTable{db: db}
}

func (c *CountryTable) UpsertCountry(ctx context.Context, alpha2Code, englishShortName string) error {
	_, err := c.db.ExecContext(ctx, `
INSERT INTO countries (alpha_2_code, english_short_name)
    VALUES (?, ?)
ON CONFLICT (alpha_2_code)
    DO UPDATE SET
        english_short_name = excluded.english_short_name;
    `, alpha2Code, englishShortName)

	if err != nil {
		return errors.Wrap(err, "db: upsert country")
	}

	return nil
}

func (c *CountryTable) UpdateDetail(ctx context.Context, detail extract.Detail) error {
	_, err := c.db.NamedExecContext(ctx, `
UPDATE countries SET 
      short_name            = :short_name
    , short_name_lower_case = :short_name_lower_case
    , full_name             = :full_name
    , alpha_3_code          = :alpha_3_code
    , numeric_code          = :numeric_code
    , remarks               = :remarks
    , independent           = :independent
    , territory_name        = :territory_name
    , status                = :status
WHERE alpha_2_code = :alpha_2_code
;`, detail)

	return errors.Wrap(err, "db: update country detail")
}

func (c *CountryTable) getCountryEnglishShortName(ctx context.Context, alpha2Code string) (string, error) {
	var englishShortName string

	err := c.db.QueryRowxContext(ctx, `
SELECT
    english_short_name
FROM
    countries
WHERE
    alpha_2_code = ?
;
    `, alpha2Code).Scan(&englishShortName)

	if err != nil {
		return "", errors.Wrap(err, "db: get english_short_name")
	}

	return englishShortName, nil
}
