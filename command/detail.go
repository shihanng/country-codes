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
	logger           *log.Logger
	countryTable     *db.CountryTable
	languageTable    *db.LanguageTable
	subdivisionTable *db.SubdivisionTable
}

func (c *detailCommand) Help() string {
	return `
Download detail country from https://www.iso.org/obp/ui#search into local DB.
`
}

func (c *detailCommand) Run(args []string) int {
	ctx := context.Background()

	if len(args) != 1 {
		c.logger.Errorf("expected 1 argument got %v", args)
		return 1
	}

	countryCodes := strings.Split(args[0], ",")

	for _, countryCode := range countryCodes {
		logCtx := c.logger.WithFields(log.Fields{
			"country_code": countryCode,
		})
		logCtx.Info("start downloading html")

		detailHTML, err := download.DownloadCountryDetailHTML(ctx, download.DetailURL(countryCode))
		if err != nil {
			logCtx.WithError(err).Error("failed to detail information of country")
			return 1
		}

		logCtx.Info("done downloading html")

		r := strings.NewReader(detailHTML)

		detail, err := extract.ExtractDetail(r)
		if err != nil {
			logCtx.WithError(err).Error("failed to extract detail")
			return 1
		}

		logCtx.Info("done extracting detail")

		if err := c.countryTable.UpdateDetail(ctx, *detail); err != nil {
			logCtx.WithError(err).Error("failed to update country detail")
			return 1
		}

		logCtx.Info("done update country detail")

		for _, language := range detail.Languages {
			logCtxLang := logCtx.WithField("language", language.Alpha2)
			logCtxLang.Info("register language")

			lang := extract.Language{
				Alpha2: strings.ToLower(language.Alpha2),
				Alpha3: language.Alpha3,
			}

			if err := c.languageTable.UpsertLanguage(ctx, lang); err != nil {
				logCtxLang.WithError(err).Error("failed to register language to db")
				return 1
			}

			if err := c.languageTable.SetLocalShortName(ctx, countryCode, lang.Alpha2, language.LocalShortName); err != nil {
				logCtxLang.WithError(err).Error("failed to register language for country to db")
			}
		}

		for _, subdivision := range detail.Subdivisions {
			logCtxSub := logCtx.WithField("code_31662", subdivision.Code31662)
			logCtxSub.Info("register subdivision")

			subdivision.CountryCode = countryCode

			if err := c.subdivisionTable.UpsertSubdivision(ctx, subdivision); err != nil {
				logCtxSub.WithError(err).Error("failed to register subdivision for country to db")
			}
		}
	}

	return 0
}

func (c *detailCommand) Synopsis() string {
	return "Download detail of country into DB"
}
