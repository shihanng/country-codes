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
	Alpha3Code         string
	NumericCode        string
	Remarks            string
	Independent        string
	TerritoryName      string
	Status             string
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
		detail.ShortNameLowerCase = s.Find("div:nth-child(3) > div.core-view-field-value").Text()
		detail.FullName = s.Find("div:nth-child(4) > div.core-view-field-value").Text()
		detail.Alpha3Code = s.Find("div:nth-child(5) > div.core-view-field-value").Text()
		detail.NumericCode = s.Find("div:nth-child(6) > div.core-view-field-value").Text()
		detail.Remarks = s.Find("div:nth-child(7) > div.core-view-field-value").Text()
		detail.Independent = s.Find("div:nth-child(8) > div.core-view-field-value").Text()
		detail.TerritoryName = s.Find("div:nth-child(9) > div.core-view-field-value").Text()
		detail.Status = s.Find("div:nth-child(10) > div.core-view-field-value").Text()
	})

	return &detail, nil
}
