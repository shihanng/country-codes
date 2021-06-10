package db

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
)

type SubdivisionTable struct {
	db *sqlx.DB
}

func NewSubdivisionTable(db *sqlx.DB) *SubdivisionTable {
	return &SubdivisionTable{db: db}
}

func (t *SubdivisionTable) UpsertSubdivision(ctx context.Context, subdivision extract.Subdivision) error {
	_, err := t.db.NamedExecContext(ctx, `
INSERT INTO subdivisions (
      country_code
    , code_31662
    , name
    , local_variant
    , language_code
    , romanization_system
    , parent_subdivision
) VALUES (
      :country_code
    , :code_31662
    , :name
    , :local_variant
    , :language_code
    , :romanization_system
    , :parent_subdivision
) ON CONFLICT (country_code, code_31662) DO UPDATE SET
      name                = excluded.name
    , local_variant       = excluded.local_variant
    , language_code       = excluded.language_code
    , romanization_system = excluded.romanization_system
    , parent_subdivision  = excluded.parent_subdivision
;`, subdivision)
	if err != nil {
		return errors.Wrap(err, "db: upsert subdivision")
	}

	return nil
}
