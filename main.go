package main

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/apex/log"
	apexcli "github.com/apex/log/handlers/cli"
	"github.com/golang-migrate/migrate/v4"
	"github.com/mitchellh/cli"
	"github.com/shihanng/country-codes/command"
	"github.com/shihanng/country-codes/db"
)

const dbPath = "country_code.db"

func main() {
	ctx := context.Background()

	fs := flag.NewFlagSet("", flag.ExitOnError)
	debug := fs.Bool("debug", false, "show debug log")

	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	logLevel := log.ErrorLevel
	if *debug {
		logLevel = log.DebugLevel
	}

	logger := log.Logger{
		Level:   logLevel,
		Handler: apexcli.New(os.Stderr),
	}

	dbInstance, err := db.NewDB(ctx, dbPath)
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("no migration applied")
		} else {
			logger.WithError(err).Error("failed create new db instance")
			return
		}
	}

	countryTable := db.NewCountryTable(dbInstance)

	logger.Info("done preparing db")

	factory := command.Factory{
		Logger:       logger,
		CountryTable: countryTable,
	}

	c := cli.NewCLI("country-codes", "0.0.0")
	c.Args = fs.Args()
	c.Commands = map[string]cli.CommandFactory{
		"update list": factory.ListCommand,
	}

	exitStatus, err := c.Run()
	if err != nil {
		logger.WithError(err).Error("failed running command")
	}

	os.Exit(exitStatus)
}
