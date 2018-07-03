package scraping

import (
	"fmt"
	"log"
	"strconv"
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
	/*add more digital currencys*/
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
	/*add more currencys*/
}

// 1 if you want to get digital currencys, or 0 if you want to get "normal" currencys
func links(currency string, typeCurrency int) (link string, err bool) {
	if typeCurrency == 1 {
		link = digitalCurrencys[currency]
		if link == "" {
			return "", false
		}
		return link, true
	} else if typeCurrency == 0 {
		link = currencys[currency]
		if link == "" {
			return "", false
		}
		return link, true
	}

	return "", false
}

// GetCurrrency get determinate currency
func GetCurrrency(currency string, typeCurrency int) Currency {
	curr, erro := links(currency, typeCurrency)

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

//GetCurrrencyChannel to implements goroutines
func GetCurrrencyChannel(currency string, typeCurrency int, ch chan Currency) {
	ch <- GetCurrrency(currency, typeCurrency)
}

// GetAllCurrency get all currencys
func GetAllCurrency(curr *[]Currency, typeCurrency int) {
	ch := make(chan Currency)
	if typeCurrency == 1 {
		for key := range digitalCurrencys {
			go GetCurrrencyChannel(key, typeCurrency, ch)
		}
		for range digitalCurrencys {
			c := <-ch
			*curr = append(*curr, c)
		}
	} else if typeCurrency == 0 {
		for key := range currencys {
			go GetCurrrencyChannel(key, typeCurrency, ch)
		}

		for range currencys {
			c := <-ch
			*curr = append(*curr, c)
		}
	} else {
		fmt.Println("Invalid parameter " + strconv.Itoa(typeCurrency))
		*curr = append(*curr, Currency{})
	}

}
