package currency

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/msalcantara/gocurrency/errors"
)

var traditionalCurrencys = map[string]string{
	"dolar":             "https://dolarhoje.com/",
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
}

type traditional map[string]string

func (t *traditional) GetAll() ([]Currency, error) {
	curr := []Currency{}
	for key := range traditionalCurrencys {
		c, err := t.GetCurrrency(key)
		if err != nil {
			return nil, err
		}
		curr = append(curr, c)
	}
	return curr, nil
}

func (t *traditional) GetCurrrency(currency string) (Currency, error) {
	curr, err := t.links(currency)

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

func (t *traditional) links(currency string) (link string, err error) {
	link = traditionalCurrencys[currency]
	if link == "" {
		return "", errors.New(fmt.Sprintf("invalid type of currency %s", currency))
	}
	return link, nil
}
