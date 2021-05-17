package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/chromedp/chromedp"
)

const (
	entryPoint           = `https://www.iso.org/obp/ui#home`
	countryCodesSelector = `#obpui-105541713 > div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div:nth-child(2) > div > div.v-slot.v-slot-xmltype > div > span:nth-child(7) > label`
	searchSelector       = `#obpui-105541713 > div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div:nth-child(2) > div > div.v-slot.v-slot-global-search.v-slot-light.v-slot-home-search > div > div.v-panel-content.v-panel-content-global-search.v-panel-content-light.v-panel-content-home-search.v-scrollable > div > div > div.v-slot.v-slot-go > div > span > span`
	tableSelector        = `#obpui-105541713 > div > div.v-customcomponent.v-widget.v-has-width.v-has-height > div > div > div:nth-child(2) > div > div > div.v-tabsheet-content.v-tabsheet-content-header > div > div > div > div > div > div.v-slot.v-slot-borderless > div > div.v-panel-content.v-panel-content-borderless.v-scrollable > div > div > div.v-slot.v-slot-search-result-layout > div > div:nth-child(2) > div.v-grid.v-widget.country-code.v-grid-country-code.v-has-width > div.v-grid-tablewrapper > table`
	outputFilename       = `country_codes.html`
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var example string
	if _, err := chromedp.RunResponse(ctx,
		chromedp.Navigate(entryPoint),
		chromedp.Click(countryCodesSelector, chromedp.NodeVisible),
		chromedp.Click(searchSelector, chromedp.NodeVisible),
		chromedp.OuterHTML(tableSelector, &example),
	); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(outputFilename, []byte(example), 0644); err != nil {
		log.Fatal(err)
	}
}
