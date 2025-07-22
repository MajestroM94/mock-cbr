package service

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

var currencyList = []struct {
	ID       string
	NumCode  string
	CharCode string
	Name     string
}{
	{"R01235", "840", "USD", "US Dollar"},
	{"R01239", "978", "EUR", "Euro"},
	{"R01375", "392", "RUB", "Russian ruble"},
	{"R01035", "826", "GBP", "Pound Sterling"},
	{"R01355", "372", "JPY", "Japanese Yen"},
}

type ValuteData struct {
	ID       string
	NumCode  string
	CharCode string
	Nominal  int
	Name     string
	Value    string
}

func GenerateRandomXML(date string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	valutes := make([]Valute, 0, len(currencyList))
	for _, c := range currencyList {
		valutes = append(valutes, Valute{
			ID:       c.ID,
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Nominal:  1,
			Name:     c.Name,
			Value:    fmt.Sprintf("%.2f", 60+rand.Float64()*50),
		})
	}

	valCurs := ValCurs{
		Date:    date,
		Name:    "Foreign Currency Market",
		Valutes: valutes,
	}

	result, err := xml.MarshalIndent(valCurs, "", "  ")
	if err != nil {
		return "", err
	}

	return xml.Header + string(result), nil
}

func GenerateRandomData(date string) (string, []ValuteData, error) {
	rand.Seed(time.Now().UnixNano())

	valutes := make([]Valute, 0, len(currencyList))
	valuteData := make([]ValuteData, 0, len(currencyList))

	for _, c := range currencyList {
		val := fmt.Sprintf("%.2f", 70+rand.Float64()*40)

		valutes = append(valutes, Valute{
			ID:       c.ID,
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Nominal:  1,
			Name:     c.Name,
			Value:    val,
		})

		valuteData = append(valuteData, ValuteData{
			ID:       c.ID,
			NumCode:  c.NumCode,
			CharCode: c.CharCode,
			Nominal:  1,
			Name:     c.Name,
			Value:    val,
		})
	}

	valCurs := ValCurs{
		Date:    date,
		Name:    "Foreign Currency Market",
		Valutes: valutes,
	}

	xmlOutput, err := xml.MarshalIndent(valCurs, "", "  ")
	if err != nil {
		return "", nil, err
	}

	return xml.Header + string(xmlOutput), valuteData, nil
}
