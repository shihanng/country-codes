package command

import (
	"context"
	"strings"

	"github.com/apex/log"
	"github.com/shihanng/country-codes/db"
	"github.com/shihanng/country-codes/download"
	"github.com/shihanng/country-codes/extract"
)

type listCommand struct {
	logger log.Logger
	table  *db.CountryTable
}

func (c *listCommand) Help() string {
	return `
Download list of countries with Alpha-2 code and English short name
from https://www.iso.org/obp/ui#search into local DB.
`
}

func (c *listCommand) Run(_ []string) int {
	ctx := context.Background()

	listHTML, err := download.DownloadCountryListHTML(ctx, download.URL)
	if err != nil {
		c.logger.WithError(err).Error("failed to download main list of countries")
		return 1
	}

	c.logger.Info("done downloading html")

	r := strings.NewReader(listHTML)

	codes, err := extract.ExtractAlpha2Code(r)
	if err != nil {
		c.logger.WithError(err).Error("failed to extract Alpha-2 codes")
		return 1
	}

	c.logger.Info("done extracting Alpha-2 codes")

	for _, code := range codes {
		if err := c.table.UpsertCountry(ctx, code.Code, code.EnglishShortName); err != nil {
			c.logger.WithError(err).Error("failed to register country to db")
			return 1
		}
	}

	return 0
}

func (c *listCommand) Synopsis() string {
	return "Download list of countries into DB"
}
