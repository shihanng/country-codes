package db

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
)

type LanguageTable struct {
	db *sqlx.DB
}

func NewLanguageTable(db *sqlx.DB) *LanguageTable {
	return &LanguageTable{db: db}
}

func (t *LanguageTable) UpsertLanguage(ctx context.Context, lang extract.Language) error {
	_, err := t.db.NamedExecContext(ctx, `
INSERT INTO languages (
      alpha_2_code
    , alpha_3_code
) VALUES (
      :alpha_2_code
    , :alpha_3_code
) ON CONFLICT (alpha_2_code) DO UPDATE SET
      alpha_3_code = excluded.alpha_3_code
;`, lang)
	if err != nil {
		return errors.Wrap(err, "db: upsert language")
	}

	return nil
}

func (t *LanguageTable) getLanguage(ctx context.Context, alpha2Code string) (*extract.Language, error) {
	var lang extract.Language

	err := t.db.QueryRowxContext(ctx, `
SELECT * FROM languages WHERE alpha_2_code = ?
;`, alpha2Code).StructScan(&lang)

	if err != nil {
		return nil, errors.Wrap(err, "db: get language")
	}

	return &lang, nil
}
