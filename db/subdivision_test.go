package db

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SubdivisionTableSuite struct {
	suite.Suite
	db               *sqlx.DB
	subdivisionTable *SubdivisionTable
	countryTable     *CountryTable
	languageTable    *LanguageTable
}

func (s *SubdivisionTableSuite) SetupTest() {
	ctx := context.Background()

	db, err := NewDB(ctx, Memory)
	s.Require().NoError(err)

	s.db = db

	s.subdivisionTable = NewSubdivisionTable(s.db)
	s.countryTable = NewCountryTable(s.db)
	s.languageTable = NewLanguageTable(s.db)

	s.Require().NoError(s.countryTable.UpsertCountry(ctx, "BR", "Brazil"))
	s.Require().NoError(s.languageTable.UpsertLanguage(ctx, extract.Language{
		Alpha2: "pt",
		Alpha3: "por",
	}))
}

func (s *SubdivisionTableSuite) TestSubdivisionTable_UpsertSubdivision() {
	type args struct {
		lang extract.Subdivision
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "CountryCode must be known",
			args: args{
				lang: extract.Subdivision{
					CountryCode: "IN",
				},
			},
			assertion: assert.Error,
		},
		{
			name: "OK",
			args: args{
				lang: extract.Subdivision{
					CountryCode:  "BR",
					Category:     "State",
					Code31662:    "BR-AC",
					Name:         "Acre",
					LocalVariant: "",
					LanguageCode: "pt",
				},
			},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := s.subdivisionTable.UpsertSubdivision(context.Background(), tt.args.lang)
			tt.assertion(t, err)
		})
	}
}

func TestSubdivisionTableSuite(t *testing.T) {
	suite.Run(t, &SubdivisionTableSuite{})
}
