package tests

import (
	"testing"

	"github.com/matheussilva97/currency-quote-api/scraping"
)

func TestGetCurrrencyDolar(t *testing.T) {
	curr := scraping.GetCurrrency("dolar", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil")
	}

	t.Logf("[PASSED] - value of dolar is " + curr.Value)
}

func TestGetCurrrencyDolarAustraliano(t *testing.T) {
	curr := scraping.GetCurrrency("dolar australiano", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar australiano'")
	}

	t.Logf("[PASSED] - value of dolar australiano is " + curr.Value)
}

func TestGetCurrrencyDolarCanadense(t *testing.T) {
	curr := scraping.GetCurrrency("dolar canadense", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar canadense'")
	}

	t.Logf("[PASSED] - value of dolar canadense is " + curr.Value)
}

func TestGetCurrrencyDolarNeozelandes(t *testing.T) {
	curr := scraping.GetCurrrency("dolar neozelandes", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar neozelandes'")
	}

	t.Logf("[PASSED] - value of dolar neozelandes is " + curr.Value)
}

func TestGetCurrrencyIene(t *testing.T) {
	curr := scraping.GetCurrrency("iene", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'iene'")
	}

	t.Logf("[PASSED] - value of iene is " + curr.Value)
}
func TestGetCurrrencyEuro(t *testing.T) {
	curr := scraping.GetCurrrency("euro", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'euro'")
	}

	t.Logf("[PASSED] - value of euro is " + curr.Value)
}

func TestGetCurrrencyLibra(t *testing.T) {
	curr := scraping.GetCurrrency("libra", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'libra'")
	}

	t.Logf("[PASSED] - value of libra is " + curr.Value)
}

func TestGetCurrrencyOuro(t *testing.T) {
	curr := scraping.GetCurrrency("ouro", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'ouro'")
	}

	t.Logf("[PASSED] - value of ouro is " + curr.Value)
}

func TestGetCurrrencyPesoArgentino(t *testing.T) {
	curr := scraping.GetCurrrency("peso argentino", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso argentino'")
	}

	t.Logf("[PASSED] - value of peso argentino is " + curr.Value)
}

func TestGetCurrrencyPesoChileno(t *testing.T) {
	curr := scraping.GetCurrrency("peso chileno", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso chileno'")
	}

	t.Logf("[PASSED] - value of peso chileno is " + curr.Value)
}

func TestGetCurrrencyPesoUruguaio(t *testing.T) {
	curr := scraping.GetCurrrency("peso uruguaio", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso uruguaio'")
	}

	t.Logf("[PASSED] - value of peso uruguaio is " + curr.Value)
}

func TestGetCurrrencyWonSuloreano(t *testing.T) {
	curr := scraping.GetCurrrency("won sul-coreano", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'won sul-coreano'")
	}

	t.Logf("[PASSED] - value of won sul-coreano is " + curr.Value)
}

func TestGetCurrrencyYuan(t *testing.T) {
	curr := scraping.GetCurrrency("yuan", 0)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'yuan'")
	}

	t.Logf("[PASSED] - value of yuan is " + curr.Value)
}

func TestGetAllCurrrencyt(t *testing.T) {
	currency := []scraping.Currency{}
	scraping.GetAllCurrency(&currency, 0)

	if currency == nil {
		t.Fatalf("[FAILED] - Function retun nil to get all currencys")
	}
}

/* === Digital currency tests ===*/

func TestGetAeternity(t *testing.T) {
	curr := scraping.GetCurrrency("aeternity", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'aeternity'")
	}

	t.Logf("[PASSED] - value of aeternity is " + curr.Value)
}

func TestGetBasicAttentionToken(t *testing.T) {
	curr := scraping.GetCurrrency("basic attention token", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'basic attention token'")
	}

	t.Logf("[PASSED] - value of basic attention token is " + curr.Value)
}
func TestGetBitcoin(t *testing.T) {
	curr := scraping.GetCurrrency("bitcoin", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin'")
	}

	t.Logf("[PASSED] - value of bitcoin is " + curr.Value)
}
func TestGetBitcoinCash(t *testing.T) {
	curr := scraping.GetCurrrency("bitcoin cash", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin cash'")
	}

	t.Logf("[PASSED] - value of bitcoin cash is " + curr.Value)
}
func TestGetbitcoinGold(t *testing.T) {
	curr := scraping.GetCurrrency("bitcoin gold", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin gold'")
	}

	t.Logf("[PASSED] - value of bitcoin gold is " + curr.Value)
}
func TestGetBitshares(t *testing.T) {
	curr := scraping.GetCurrrency("bitshares", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitshares'")
	}

	t.Logf("[PASSED] - value of bitshares is " + curr.Value)
}
func TestGetBytecoin(t *testing.T) {
	curr := scraping.GetCurrrency("bytecoin", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bytecoin'")
	}

	t.Logf("[PASSED] - value of bytecoin is " + curr.Value)
}
func TestGetBytom(t *testing.T) {
	curr := scraping.GetCurrrency("bytom", 1)

	if curr == (scraping.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bytom'")
	}

	t.Logf("[PASSED] - value of bytom is " + curr.Value)
}

func TestGetAllDigitalCurrrencyt(t *testing.T) {
	currency := []scraping.Currency{}
	scraping.GetAllCurrency(&currency, 1)

	if currency == nil {
		t.Fatalf("[FAILED] - Function retun nil to get all currencys")
	}
}
