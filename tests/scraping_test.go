package tests

import (
	"github.com/matheussilva97/currency-quote-api/scraping"
	"testing"
)

func TestGetCurrrencyDolar(t *testing.T) {
	curr := scraping.GetCurrrency("dolar")

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil")
	}

	t.Logf("[PASSED] - value of dolar is " + curr.Value)
}

func TestGetCurrrencyEuro(t *testing.T) {
	curr := scraping.GetCurrrency("euro")

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil")
	}

	t.Logf("[PASSED] - value of euro is " + curr.Value)
}

func TestGetCurrrencyLibra(t *testing.T) {
	curr := scraping.GetCurrrency("libra")

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'libra'")
	}

	t.Logf("[PASSED] - value of libra is " + curr.Value)
}

func TestGetAllCurrrencyt(t *testing.T) {
	done := make(chan bool)
	currency := []scraping.Currency{}
	go scraping.GetAllCurrency(&currency, done)
	<-done

	if currency == nil {
		t.Fatalf("[FAILED] - Function retun nil to param 'libra'")
	}
}

/* === Digital currency tests ===*/
