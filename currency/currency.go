package currency

import (
	"errors"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Currency struct {
	Name             string `json:"currency_name"`
	Value            string `json:"currency_value"`
	VerificationDate string `json:"verification_date"`
}

type ICurrency interface {
	//GetAll return currency quote of all currencys
	GetAll() ([]Currency, error)

	//GetOne return currency quote of specific currency
	GetOne(string) (Currency, error)
}

type CurrencyImp struct {
	Digital     ICurrency
	Traditional ICurrency
}

type urlScraping map[string]string

func NewCurrency() *CurrencyImp {
	return &CurrencyImp{
		Digital: digitalCurrency{
			urls: map[string]string{
				"aeternity":             "https://dolarhoje.com/aeternity-hoje/",
				"basic attention token": "https://dolarhoje.com/basic-attention-token-hoje/",
				"bitcoin":               "https://dolarhoje.com/bitcoin-hoje/",
				"bitcoin cash":          "https://dolarhoje.com/bitcoin-cash-hoje/",
				"bitcoin gold":          "https://dolarhoje.com/bitcoin-gold-hoje/",
				"bitshares":             "https://dolarhoje.com/bitshares-hoje/",
				"bytecoin":              "https://dolarhoje.com/bytecoin-hoje/",
				"bytom":                 "https://dolarhoje.com/bytom-hoje/",
			},
		},
		Traditional: traditionalCurrency{
			urls: map[string]string{
				"dolar":             "https://ddolarhoje.com/",
				"dolar australiano": "https://dolarhoje.com/dolar-australiano-hoje/",
				"dolar canadense":   "https://dolarhoje.com/dolar-canadense-hoje/",
				"dolar neozelandes": "https://dolarhoje.com/dolar-neozelandes-hoje/",
				"iene":              "https://dolarhoje.com/iene/",
				"euro":              "https://dolarhoje.com/euro/",
				"libra":             "https://dolarhoje.com/libra-hoje/",
				"ouro":              "https://dolarhoje.com/ouro-hoje/",
				"peso argentino":    "https://dolarhoje.com/peso-argentino/",
				"peso chileno":      "https://dolarhoje.com/peso-chileno/",
				"peso uruguaio":     "https://dolarhoje.com/peso-uruguaio/",
				"won sul-coreano":   "https://dolarhoje.com/won-sul-coreano-hoje/",
				"yuan":              "https://dolarhoje.com/yuan-hoje/",
			},
		},
	}
}

type digitalCurrency struct {
	urls urlScraping
}

type traditionalCurrency struct {
	urls urlScraping
}

func (tc traditionalCurrency) GetAll() ([]Currency, error) {
	curr := []Currency{}
	for key := range tc.urls {
		c, err := tc.GetOne(key)
		if err != nil {
			return nil, err
		}
		curr = append(curr, c)
	}
	return curr, nil
}

func (tc traditionalCurrency) GetOne(c string) (Currency, error) {
	curr, err := tc.urls.GetURL(c)
	if err != nil {
		return Currency{}, err
	}
	doc, err := goquery.NewDocument(curr)
	if err != nil {
		return Currency{}, err
	}
	val, _ := doc.Find("input#nacional").Attr("value")
	if val == "" {
		return Currency{}, errors.New("Cannot find value of currency in the page")
	}
	return Currency{c, val, time.Now().Local().Format("2006-01-02")}, nil
}

func (dc digitalCurrency) GetAll() ([]Currency, error) {
	curr := []Currency{}
	for key := range dc.urls {
		c, err := dc.GetOne(key)
		if err != nil {
			return nil, err
		}
		curr = append(curr, c)
	}
	return curr, nil
}

func (dc digitalCurrency) GetOne(c string) (Currency, error) {
	curr, err := dc.urls.GetURL(c)

	if err != nil {
		return Currency{}, err
	}
	doc, err := goquery.NewDocument(curr)
	if err != nil {
		return Currency{}, err
	}
	val, _ := doc.Find("input#nacional").Attr("value")
	if val == "" {
		return Currency{}, errors.New("Cannot find value of currency in the page")
	}
	return Currency{c, val, time.Now().Local().Format("2006-01-02")}, nil
}

func (u urlScraping) GetURL(c string) (string, error) {
	url, ok := u[c]
	if !ok {
		return "", fmt.Errorf("%s not found", c)
	}
	return url, nil
}
