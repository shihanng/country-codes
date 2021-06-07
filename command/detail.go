package command

import (
	"context"
	"strings"

	"github.com/apex/log"
	"github.com/shihanng/country-codes/db"
	"github.com/shihanng/country-codes/download"
	"github.com/shihanng/country-codes/extract"
)

type detailCommand struct {
	logger        log.Logger
	countryTable  *db.CountryTable
	languageTable *db.LanguageTable
}

func (c *detailCommand) Help() string {
	return `
Download detail country from https://www.iso.org/obp/ui#search into local DB.
`
}

func (c *detailCommand) Run(args []string) int {
	ctx := context.Background()

	detailHTML, err := download.DownloadCountryDetailHTML(ctx, download.DetailURL("CH"))
	if err != nil {
		c.logger.WithError(err).Error("failed to detail information of country")
		return 1
	}

	c.logger.Info("done downloading html")

	r := strings.NewReader(detailHTML)

	detail, err := extract.ExtractDetail(r)
	if err != nil {
		c.logger.WithError(err).Error("failed to extract detail")
		return 1
	}

	c.logger.Info("done extracting detail")

	for _, language := range detail.Languages {
		lang := extract.Language{
			Alpha2: strings.ToLower(language.Alpha2),
			Alpha3: language.Alpha3,
		}

		if err := c.languageTable.UpsertLanguage(ctx, lang); err != nil {
			c.logger.WithError(err).Error("failed to register language to db")
			return 1
		}
	}

	return 0
}

func (c *detailCommand) Synopsis() string {
	return "Download detail of country into DB"
}
