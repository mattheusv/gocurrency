package tests

import (
	"net/http"
	"testing"
)

func TestGetAllTraditionalCurrencysHeroku(t *testing.T) {
	response, err := http.Get("https://gocurrency.herokuapp.com/api/currencys")
	if err != nil {
		t.Errorf("could not make request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", response.StatusCode, http.StatusOK)
	}
}

func TestGetOneTraditionalCurrencysHeroku(t *testing.T) {
	response, err := http.Get("https://gocurrency.herokuapp.com/api/currencys/dolar")
	if err != nil {
		t.Errorf("could not make request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", response.StatusCode, http.StatusOK)
	}
}

func TestGetAllDigitalCurrencysHeroku(t *testing.T) {
	response, err := http.Get("https://gocurrency.herokuapp.com/api/digital/currencys")
	if err != nil {
		t.Errorf("could not make request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", response.StatusCode, http.StatusOK)
	}
}

func TestGetOneDigitalCurrencysHeroku(t *testing.T) {
	response, err := http.Get("https://gocurrency.herokuapp.com/api/digital/currencys/bitcoin")
	if err != nil {
		t.Errorf("could not make request: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", response.StatusCode, http.StatusOK)
	}
}
