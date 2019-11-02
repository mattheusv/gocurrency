package currency

import "testing"

func TestGetURL(t *testing.T) {
	u := urlScraping{"dolar": "http://localhost:9090"}
	url, err := u.GetURL("dolar")
	if err != nil {
		t.Fatal(err)
	}
	if url != "http://localhost:9090" {
		t.Errorf("invalid url: %s", url)
	}
}

func TestGetURLError(t *testing.T) {
	u := urlScraping{"dolar": "http://localhost:9090"}
	_, err := u.GetURL("libra")
	if err == nil {
		t.Errorf("expected error invalid currency")
	}
}
