package extract

import (
	"io"
	"strings"

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
	Languages          []AdministrativeLanguage
	Subdivisions       []Subdivision
}

type AdministrativeLanguage struct {
	Alpha2         string
	Alpha3         string
	LocalShortName string
}

type Subdivision struct {
	Category           string
	Code31662          string
	Name               string
	LocalVariant       string
	LanguageCode       string
	RomanizationSystem string
	ParentSubdivision  string
}

func ExtractDetail(r io.Reader) (*Detail, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var detail Detail
	doc.Find("div.core-view-summary").Each(func(_ int, s *goquery.Selection) {
		detail.Alpha2Code = strings.TrimSpace(s.Find("div:nth-child(1) > div.core-view-field-value").Text())
		detail.ShortName = strings.TrimSpace(s.Find("div:nth-child(2) > div.core-view-field-value").Text())
		detail.ShortNameLowerCase = strings.TrimSpace(s.Find("div:nth-child(3) > div.core-view-field-value").Text())
		detail.FullName = strings.TrimSpace(s.Find("div:nth-child(4) > div.core-view-field-value").Text())
		detail.Alpha3Code = strings.TrimSpace(s.Find("div:nth-child(5) > div.core-view-field-value").Text())
		detail.NumericCode = strings.TrimSpace(s.Find("div:nth-child(6) > div.core-view-field-value").Text())
		detail.Remarks = strings.TrimSpace(s.Find("div:nth-child(7) > div.core-view-field-value").Text())
		detail.Independent = strings.TrimSpace(s.Find("div:nth-child(8) > div.core-view-field-value").Text())
		detail.TerritoryName = strings.TrimSpace(s.Find("div:nth-child(9) > div.core-view-field-value").Text())
		detail.Status = strings.TrimSpace(s.Find("div:nth-child(10) > div.core-view-field-value").Text())
	})

	doc.Find("#country-additional-info > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		var al AdministrativeLanguage
		al.Alpha2 = strings.TrimSpace(s.Find("#country-additional-info > table > tbody > tr > td:nth-child(1)").Text())
		al.Alpha3 = strings.TrimSpace(s.Find("#country-additional-info > table > tbody > tr > td:nth-child(2)").Text())
		al.LocalShortName = strings.TrimSpace(s.Find("#country-additional-info > table > tbody > tr > td:nth-child(3)").Text())

		detail.Languages = append(detail.Languages, al)
	})

	doc.Find("#subdivision > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		var sub Subdivision
		sub.Category = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(1)").Text())
		sub.Code31662 = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(2)").Text())
		sub.Name = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(3)").Text())
		sub.LocalVariant = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(4)").Text())
		sub.LanguageCode = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(5)").Text())
		sub.RomanizationSystem = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(6)").Text())
		sub.ParentSubdivision = strings.TrimSpace(s.Find("#subdivision > tbody > tr > td:nth-child(7)").Text())

		detail.Subdivisions = append(detail.Subdivisions, sub)
	})

	return &detail, nil
}
