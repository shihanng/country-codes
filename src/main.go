package main

import (
	"context"
	"encoding/csv"
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/cockroachdb/errors"
	"github.com/shihanng/country-codes/src/download"
	"github.com/shihanng/country-codes/src/extract"
)

func main() {
	logger := log.Logger{
		Level:   log.InfoLevel,
		Handler: cli.New(os.Stderr),
	}

	listHTML, err := download.DownloadCountryListHTML(context.Background(), download.URL)
	if err != nil {
		logger.WithError(err).Error("failed to download main list of countries")
	}

	logger.Info("done downloading html")

	r := strings.NewReader(listHTML)

	codes, err := extract.ExtractAlpha2Code(r)
	if err != nil {
		logger.WithError(err).Error("failed to extract Alpha-2 codes")
	}

	logger.Info("done extracting Alpha-2 codes")

	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"code", "english_short_name"}); err != nil {
		logger.WithError(errors.Wrap(err, "write csv header")).Error("failed to write csv header")
	}

	for _, code := range codes {
		if err := w.Write([]string{code.Code, code.EnglishShortName}); err != nil {
			logger.WithError(errors.Wrap(err, "write csv row")).Error("failed to write csv row")
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		logger.WithError(errors.Wrap(err, "flush csv writer")).Error("failed flush csv")
	}
}
