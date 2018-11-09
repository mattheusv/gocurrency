package currency

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/msalcantara/gocurrency/errors"
)

var digitalCurrencys = map[string]string{
	"aeternity":             "https://dolarhoje.com/aeternity-hoje/",
	"basic attention token": "https://dolarhoje.com/basic-attention-token-hoje/",
	"bitcoin":               "https://dolarhoje.com/bitcoin-hoje/",
	"bitcoin cash":          "https://dolarhoje.com/bitcoin-cash-hoje/",
	"bitcoin gold":          "https://dolarhoje.com/bitcoin-gold-hoje/",
	"bitshares":             "https://dolarhoje.com/bitshares-hoje/",
	"bytecoin":              "https://dolarhoje.com/bytecoin-hoje/",
	"bytom":                 "https://dolarhoje.com/bytom-hoje/",
}

type digital map[string]string

func (d *digital) GetAll() ([]Currency, error) {
	curr := []Currency{}
	for key := range digitalCurrencys {
		c, err := d.GetCurrrency(key)
		if err != nil {
			return nil, err
		}
		curr = append(curr, c)
	}
	return curr, nil
}

func (d *digital) GetCurrrency(currency string) (Currency, error) {
	curr, err := d.links(currency)

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
	return Currency{currency, "R$ " + val, time.Now().Local().Format("2006-01-02")}, nil
}

func (d *digital) links(currency string) (link string, err error) {
	link = digitalCurrencys[currency]
	if link == "" {
		return "", errors.New(fmt.Sprintf("invalid type of currency %s", currency))
	}
	return link, nil
}
