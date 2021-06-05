package extract

import (
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Alpha2Code struct {
	EnglishShortName string
	Code             string
}

func ExtractAlpha2Code(r io.Reader) ([]Alpha2Code, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var codes []Alpha2Code
	doc.Find("tr").Each(func(_ int, s *goquery.Selection) {
		country := s.Find("button").Text()
		alpha2Code := s.Find("td:nth-child(3)").Text()

		if country == "" || alpha2Code == "" {
			return
		}

		codes = append(codes, Alpha2Code{
			EnglishShortName: country,
			Code:             alpha2Code,
		})
	})

	return codes, nil
}

type Detail struct {
	Alpha2Code         string
	ShortName          string
	ShortNameLowerCase string
	FullName           string
}

func ExtractDetail(r io.Reader) (*Detail, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var detail Detail
	doc.Find("div.core-view-summary").Each(func(_ int, s *goquery.Selection) {
		detail.Alpha2Code = s.Find("div:nth-child(1) > div.core-view-field-value").Text()
		detail.ShortName = s.Find("div:nth-child(2) > div.core-view-field-value").Text()
	})

	return &detail, nil
}
