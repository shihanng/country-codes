package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/apex/log"
	apexcli "github.com/apex/log/handlers/cli"
	"github.com/golang-migrate/migrate/v4"
	"github.com/mitchellh/cli"
	"github.com/shihanng/country-codes/command"
	"github.com/shihanng/country-codes/db"
)

const dbPath = "country_code.db?_foreign_keys=1"

func main() {
	ctx := context.Background()

	fs := flag.NewFlagSet("", flag.ExitOnError)

	logger := log.Logger{
		Level:   log.ErrorLevel,
		Handler: apexcli.New(os.Stderr),
	}

	factory := command.Factory{
		Logger: &logger,
	}

	c := cli.NewCLI("country-codes", "0.0.0")

	c.Commands = map[string]cli.CommandFactory{
		"update list":   factory.ListCommand,
		"update detail": factory.DetailCommand,
	}

	fs.Usage = func() {
		fmt.Println(c.HelpFunc(c.Commands))
		fs.PrintDefaults()
	}

	debug := fs.Bool("debug", false, "show debug log")
	version := fs.Bool("version", false, "print version")

	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	if *debug {
		logger.Level = log.DebugLevel
	}

	if *version {
		fmt.Println(c.Version)
		os.Exit(0)
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

	factory.CountryTable = db.NewCountryTable(dbInstance)
	factory.LanguageTable = db.NewLanguageTable(dbInstance)

	logger.Info("done preparing db")

	c.Args = fs.Args()

	exitStatus, err := c.Run()
	if err != nil {
		logger.WithError(err).Error("failed running command")
	}

	os.Exit(exitStatus)
}
