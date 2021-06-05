package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/shihanng/country-codes/src/download"
	"github.com/shihanng/country-codes/src/extract"
)

func main() {
	listHTML, err := download.DownloadCountryListHTML(context.Background(), download.URL)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(listHTML)

	codes, err := extract.ExtractAlpha2Code(r)
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{"code", "english_short_name"}); err != nil {
		log.Fatalln("error writing record to csv:", err)
	}

	for _, code := range codes {
		if err := w.Write([]string{code.Code, code.EnglishShortName}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
