package download

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/cockroachdb/errors"
)

const (
	URL = `https://www.iso.org/obp/ui#home`

	countryCodesSelector   = `div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div:nth-child(2) > div > div.v-slot.v-slot-xmltype > div > span:nth-child(7) > label`
	searchSelector         = `div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div:nth-child(2) > div > div.v-slot.v-slot-global-search.v-slot-light.v-slot-home-search > div > div.v-panel-content.v-panel-content-global-search.v-panel-content-light.v-panel-content-home-search.v-scrollable > div > div > div.v-slot.v-slot-go > div > span > span`
	tableSelector          = `div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div.v-slot.v-slot-borderless > div > div.v-panel-content.v-panel-content-borderless.v-scrollable > div > div > div.v-slot.v-slot-search-result-layout > div > div:nth-child(2) > div.v-grid.v-widget.v-has-width.country-code.v-grid-country-code > div.v-grid-tablewrapper > table`
	resultsPerPageSelector = `div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div.v-slot.v-slot-search-header > div > div:nth-child(5) > div:nth-child(3) > div > select`
	rowSelector            = `div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div.v-slot.v-slot-borderless > div > div.v-panel-content.v-panel-content-borderless.v-scrollable > div > div > div.v-slot.v-slot-search-result-layout > div > div:nth-child(2) > div.v-grid.v-widget.country-code.v-grid-country-code.v-has-width > div.v-grid-tablewrapper > table > tbody > tr:nth-child(242) > td:nth-child(2)`

	countryName = `Vierges`
)

func DownloadCountryListHTML(ctx context.Context, url string) (string, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	defer cancel()

	var htmlContent string

	if _, err := chromedp.RunResponse(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(countryCodesSelector, chromedp.ByQuery),
		chromedp.Click(countryCodesSelector, chromedp.NodeVisible),
		chromedp.Click(searchSelector, chromedp.NodeVisible),
		chromedp.SetValue(resultsPerPageSelector, "8"), // Show 300 results per page.
		chromedp.WaitVisible(rowSelector),
		chromedp.PollFunction(`(sel, countryName) => document.querySelector(sel).innerText.includes(countryName)`, nil,
			chromedp.WithPollingArgs(rowSelector, countryName)),
		chromedp.OuterHTML(tableSelector, &htmlContent),
	); err != nil {
		return "", errors.Wrap(err, "download: download country list")
	}

	return htmlContent, nil
}
