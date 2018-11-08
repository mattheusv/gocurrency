package tests

import (
	"testing"

	"github.com/msalcantara/currency-quote-api/currency"
)

func TestGetCurrrencyDolar(t *testing.T) {
	curr, err := currency.GetCurrrency("dolar", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil")
	}

	t.Logf("[PASSED] - value of dolar is " + curr.Value)
}

func TestGetCurrrencyDolarAustraliano(t *testing.T) {
	curr, err := currency.GetCurrrency("dolar australiano", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar australiano'")
	}

	t.Logf("[PASSED] - value of dolar australiano is " + curr.Value)
}

func TestGetCurrrencyDolarCanadense(t *testing.T) {
	curr, err := currency.GetCurrrency("dolar canadense", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar canadense'")
	}

	t.Logf("[PASSED] - value of dolar canadense is " + curr.Value)
}

func TestGetCurrrencyDolarNeozelandes(t *testing.T) {
	curr, err := currency.GetCurrrency("dolar neozelandes", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'dolar neozelandes'")
	}

	t.Logf("[PASSED] - value of dolar neozelandes is " + curr.Value)
}

func TestGetCurrrencyIene(t *testing.T) {
	curr, err := currency.GetCurrrency("iene", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'iene'")
	}

	t.Logf("[PASSED] - value of iene is " + curr.Value)
}
func TestGetCurrrencyEuro(t *testing.T) {
	curr, err := currency.GetCurrrency("euro", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'euro'")
	}

	t.Logf("[PASSED] - value of euro is " + curr.Value)
}

func TestGetCurrrencyLibra(t *testing.T) {
	curr, err := currency.GetCurrrency("libra", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'libra'")
	}

	t.Logf("[PASSED] - value of libra is " + curr.Value)
}

func TestGetCurrrencyOuro(t *testing.T) {
	curr, err := currency.GetCurrrency("ouro", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'ouro'")
	}

	t.Logf("[PASSED] - value of ouro is " + curr.Value)
}

func TestGetCurrrencyPesoArgentino(t *testing.T) {
	curr, err := currency.GetCurrrency("peso argentino", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso argentino'")
	}

	t.Logf("[PASSED] - value of peso argentino is " + curr.Value)
}

func TestGetCurrrencyPesoChileno(t *testing.T) {
	curr, err := currency.GetCurrrency("peso chileno", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso chileno'")
	}

	t.Logf("[PASSED] - value of peso chileno is " + curr.Value)
}

func TestGetCurrrencyPesoUruguaio(t *testing.T) {
	curr, err := currency.GetCurrrency("peso uruguaio", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'peso uruguaio'")
	}

	t.Logf("[PASSED] - value of peso uruguaio is " + curr.Value)
}

func TestGetCurrrencyWonSuloreano(t *testing.T) {
	curr, err := currency.GetCurrrency("won sul-coreano", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'won sul-coreano'")
	}

	t.Logf("[PASSED] - value of won sul-coreano is " + curr.Value)
}

func TestGetCurrrencyYuan(t *testing.T) {
	curr, err := currency.GetCurrrency("yuan", currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'yuan'")
	}

	t.Logf("[PASSED] - value of yuan is " + curr.Value)
}

func TestGetAllCurrrencyt(t *testing.T) {
	curr, err := currency.GetAllCurrency(currency.TRADITIONAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}
	if curr == nil {
		t.Fatalf("[FAILED] - Function retun nil to get all currencys")
	}
}

/* === Digital currency tests ===*/

func TestGetAeternity(t *testing.T) {
	curr, err := currency.GetCurrrency("aeternity", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'aeternity'")
	}

	t.Logf("[PASSED] - value of aeternity is " + curr.Value)
}

func TestGetBasicAttentionToken(t *testing.T) {
	curr, err := currency.GetCurrrency("basic attention token", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'basic attention token'")
	}

	t.Logf("[PASSED] - value of basic attention token is " + curr.Value)
}
func TestGetBitcoin(t *testing.T) {
	curr, err := currency.GetCurrrency("bitcoin", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin'")
	}

	t.Logf("[PASSED] - value of bitcoin is " + curr.Value)
}
func TestGetBitcoinCash(t *testing.T) {
	curr, err := currency.GetCurrrency("bitcoin cash", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin cash'")
	}

	t.Logf("[PASSED] - value of bitcoin cash is " + curr.Value)
}
func TestGetbitcoinGold(t *testing.T) {
	curr, err := currency.GetCurrrency("bitcoin gold", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitcoin gold'")
	}

	t.Logf("[PASSED] - value of bitcoin gold is " + curr.Value)
}
func TestGetBitshares(t *testing.T) {
	curr, err := currency.GetCurrrency("bitshares", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bitshares'")
	}

	t.Logf("[PASSED] - value of bitshares is " + curr.Value)
}
func TestGetBytecoin(t *testing.T) {
	curr, err := currency.GetCurrrency("bytecoin", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bytecoin'")
	}

	t.Logf("[PASSED] - value of bytecoin is " + curr.Value)
}

func TestGetBytom(t *testing.T) {
	curr, err := currency.GetCurrrency("bytom", currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == (currency.Currency{}) {
		t.Fatalf("[FAILED] - Function retun nil to param 'bytom'")
	}

	t.Logf("[PASSED] - value of bytom is " + curr.Value)
}

func TestGetAllDigitalCurrrencyt(t *testing.T) {
	curr, err := currency.GetAllCurrency(currency.DIGITAL)
	if err != nil {
		t.Errorf("Got error: %v", err)
	}

	if curr == nil {
		t.Fatalf("[FAILED] - Function retun nil to get all currencys")
	}
}
