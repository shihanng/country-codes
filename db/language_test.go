package db

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LanguageTableSuite struct {
	suite.Suite
	db            *sqlx.DB
	languageTable *LanguageTable
	countryTable  *CountryTable
}

func (s *LanguageTableSuite) SetupTest() {
	db, err := NewDB(context.Background(), Memory)
	s.Require().NoError(err)
	s.db = db
	s.languageTable = NewLanguageTable(s.db)
	s.countryTable = NewCountryTable(s.db)
}

func (s *LanguageTableSuite) TestLanguageTable_UpsertLanguage() {
	type args struct {
		lang extract.Language
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "alpha2Code cannot be empty string",
			args: args{
				lang: extract.Language{
					Alpha2: "",
					Alpha3: "",
				},
			},
			assertion: assert.Error,
		},
		{
			name: "englishShortName can be empty",
			args: args{
				lang: extract.Language{
					Alpha2: "pt",
					Alpha3: "por",
				},
			},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := s.languageTable.UpsertLanguage(context.Background(), tt.args.lang)
			tt.assertion(t, err)
		})
	}
}

func (s *LanguageTableSuite) TestLanguageTable_UpsertLanguage_Upsert() {
	ctx := context.Background()

	lang := extract.Language{
		Alpha2: "pt",
		Alpha3: "por",
	}

	s.NoError(s.languageTable.UpsertLanguage(ctx, extract.Language{
		Alpha2: "pt",
		Alpha3: "por",
	}))
	s.NoError(s.languageTable.UpsertLanguage(ctx, lang))

	actual, err := s.languageTable.getLanguage(ctx, "pt")
	s.NoError(err)
	s.Equal(&lang, actual)
}

func (s *LanguageTableSuite) TestLanguageTable_SetLocalShortName() {
	ctx := context.Background()

	s.Require().NoError(s.countryTable.UpsertCountry(ctx, "BR", "Brazil"))

	s.Require().NoError(s.languageTable.UpsertLanguage(ctx, extract.Language{
		Alpha2: "pt",
		Alpha3: "por",
	}))

	type args struct {
		countryCode    string
		languageCode   string
		localShortName string
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "both country and language are known",
			args: args{
				countryCode:    "BR",
				languageCode:   "pt",
				localShortName: "Brasil (o)",
			},
			assertion: assert.NoError,
		},
		{
			name: "both country and language are not known",
			args: args{
				countryCode:    "MY",
				languageCode:   "ms",
				localShortName: "Malaysia",
			},
			assertion: assert.Error,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := s.languageTable.SetLocalShortName(context.Background(),
				tt.args.countryCode,
				tt.args.languageCode,
				tt.args.localShortName)
			tt.assertion(t, err)
		})
	}
}

func TestLanguageTableSuite(t *testing.T) {
	suite.Run(t, &LanguageTableSuite{})
}
