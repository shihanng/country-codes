package db

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CountryTableSuite struct {
	suite.Suite
	db           *sqlx.DB
	countryTable *CountryTable
}

func (s *CountryTableSuite) SetupSuite() {
	db, err := NewDB(Memory)
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
			err := s.countryTable.UpsertCountry(tt.args.alpha2Code, tt.args.englishShortName)
			tt.assertion(t, err)
		})
	}
}

func TestCountryTableSuite(t *testing.T) {
	suite.Run(t, &CountryTableSuite{})
}
