package command

import (
	"bytes"
	"flag"

	"github.com/apex/log"
	"github.com/mitchellh/cli"
	"github.com/shihanng/country-codes/db"
)

type Factory struct {
	Logger        log.Logger
	CountryTable  *db.CountryTable
	LanguageTable *db.LanguageTable
}

func (f *Factory) ListCommand() (cli.Command, error) {
	b := bytes.Buffer{}

	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.SetOutput(&b)

	return &listCommand{
		logger:  f.Logger,
		table:   f.CountryTable,
		fs:      fs,
		buf:     &b,
		flagCSV: fs.Bool("csv", false, "print to screen in comma separated value"),
	}, nil
}

func (f *Factory) DetailCommand() (cli.Command, error) {
	return &detailCommand{
		logger:        f.Logger,
		countryTable:  f.CountryTable,
		languageTable: f.LanguageTable,
	}, nil
}
