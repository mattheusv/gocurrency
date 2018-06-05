package scraping

import (
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Currency struct
type Currency struct {
	Name             string `json:"currency_name"`
	Value            string `json:"currency_value"`
	DateVerification string `json:"verification_date"`
}

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

var currencys = map[string]string{
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

func links(currency string) (link string, err bool) {
	link = currencys[currency]
	if link == "" {
		return "", false
	}
	return link, true
}

// GetCurrrency get determinate currency
func GetCurrrency(currency string) Currency {
	curr, erro := links(currency)

	if erro == false {
		return Currency{}
	}

	doc, err := goquery.NewDocument(curr)
	if err != nil {
		log.Fatal(err)
		return Currency{}
	}
	val, _ := doc.Find("input#nacional").Attr("value")
	if val == "" {
		fmt.Println("[ERROR] - Cannot find value of currency in the page")
		return Currency{}
	}
	return Currency{currency, "R$ " + val, time.Now().Local().Format("2006-01-02")}
}

// GetAllCurrency get all currencys
func GetAllCurrency(curr *[]Currency, done chan bool) {
	defer func() {
		done <- true
	}()
	currentDate := time.Now().Local()
	for key := range currencys {
		c := GetCurrrency(key)
		*curr = append(*curr, Currency{key, c.Value, currentDate.Format("2006-01-02")})
	}
}
