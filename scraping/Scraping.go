package scraping

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Currency struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Date  string `json:"date"`
}

var currencys = map[string]string{
	"dolar": "https://dolarhoje.com/",
	"euro":  "https://dolarhoje.com/euro/",
	"libra": "https://dolarhoje.com/libra-hoje/",
}

func links(currency string) (link string, err bool) {
	link = currencys[currency]
	if link == "" {
		return "", false
	}
	return link, true
}

// GetCurrrency get value of determinate currency
func GetCurrrency(currency string) string {
	curr, erro := links(currency)

	if erro == false {
		return "Invalid parameter"
	}

	doc, err := goquery.NewDocument(curr)
	if err != nil {
		log.Fatal(err)
		return "Error"
	}
	val, _ := doc.Find("input#nacional").Attr("value")
	if val == "" {
		fmt.Println("[ERROR] - Cannot find value of currency in the page")
	}
	return val
}

// GetAllCurrency get all values of currencys
func GetAllCurrency(curr *[]Currency, done chan bool) {
	defer func() {
		done <- true
	}()
	for key := range currencys {
		c := Currency{key, GetCurrrency(key), "teste"}
		*curr = append(*curr, c)
	}
}
