package main

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/golang-migrate/migrate/v4"
	"github.com/shihanng/country-codes/db"
	"github.com/shihanng/country-codes/download"
	"github.com/shihanng/country-codes/extract"
)

const dbPath = "country_code.db"

func main() {
	ctx := context.Background()

	logger := log.Logger{
		Level:   log.InfoLevel,
		Handler: cli.New(os.Stderr),
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

	listHTML, err := download.DownloadCountryListHTML(ctx, download.URL)
	if err != nil {
		logger.WithError(err).Error("failed to download main list of countries")
		return
	}

	logger.Info("done downloading html")

	r := strings.NewReader(listHTML)

	codes, err := extract.ExtractAlpha2Code(r)
	if err != nil {
		logger.WithError(err).Error("failed to extract Alpha-2 codes")
		return
	}

	logger.Info("done extracting Alpha-2 codes")

	for _, code := range codes {
		if err := countryTable.UpsertCountry(ctx, code.Code, code.EnglishShortName); err != nil {
			logger.WithError(err).Error("failed to register country to db")
			return
		}
	}
}
