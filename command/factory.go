package command

import (
	"github.com/apex/log"
	"github.com/mitchellh/cli"
	"github.com/shihanng/country-codes/db"
)

type Factory struct {
	Logger       log.Logger
	CountryTable *db.CountryTable
}

func (f *Factory) ListCommand() (cli.Command, error) {
	return &listCommand{
		logger: f.Logger,
		table:  f.CountryTable,
	}, nil
}
