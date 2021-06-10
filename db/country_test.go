package db

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/shihanng/country-codes/extract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CountryTableSuite struct {
	suite.Suite
	db           *sqlx.DB
	countryTable *CountryTable
}

func (s *CountryTableSuite) SetupSuite() {
	db, err := NewDB(context.Background(), Memory)
	s.Require().NoError(err)
	s.db = db
	s.countryTable = NewCountryTable(s.db)
}

func (s *CountryTableSuite) TestCountryTable_UpsertCountry() {
	type args struct {
		alpha2Code       string
		englishShortName string
	}
	tests := []struct {
		name      string
		args      args
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "alpha2Code cannot be empty string",
			args: args{
				alpha2Code:       "",
				englishShortName: "Anguilla",
			},
			assertion: assert.Error,
		},
		{
			name: "englishShortName can be empty",
			args: args{
				alpha2Code:       "AI",
				englishShortName: "",
			},
			assertion: assert.NoError,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			err := s.countryTable.UpsertCountry(
				context.Background(),
				tt.args.alpha2Code,
				tt.args.englishShortName)
			tt.assertion(t, err)
		})
	}
}

func (s *CountryTableSuite) TestCountryTable_UpsertCountry_Upsert() {
	ctx := context.Background()
	s.NoError(s.countryTable.UpsertCountry(ctx, "MY", ""))
	s.NoError(s.countryTable.UpsertCountry(ctx, "MY", "Malaysia"))

	actual, err := s.countryTable.getCountryEnglishShortName(ctx, "MY")
	s.NoError(err)
	s.Equal("Malaysia", actual)
}

func (s *CountryTableSuite) TestCountryTable_UpdateDetail() {
	ctx := context.Background()
	s.NoError(s.countryTable.UpsertCountry(ctx, "MY", ""))

	d := extract.Detail{
		Alpha2Code:         "MY",
		ShortName:          "MALAYSIA",
		ShortNameLowerCase: "Malaysia",
	}

	s.NoError(s.countryTable.UpdateDetail(ctx, d))
}

func TestCountryTableSuite(t *testing.T) {
	suite.Run(t, &CountryTableSuite{})
}
