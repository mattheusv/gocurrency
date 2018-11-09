package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/msalcantara/gocurrency/api"
)

func TestGetAllTraditionalCurrencys(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/currencys", nil)
	if err != nil {
		t.Errorf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetCurrencys)
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", rec.Code, http.StatusOK)
	}

	errHandler := `{"data":{}}`

	if rec.Body.String() == errHandler {
		t.Errorf("handler returned empty list of traditional currencys")
	}
}

func TestGetAllDigitalCurrencys(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/digital/currencys", nil)
	if err != nil {
		t.Errorf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetDigitalCurrencys)
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("handler return wrong status code\n got: %d, want: %d", rec.Code, http.StatusOK)
	}

	errHandler := `{"data":{}}`

	if rec.Body.String() == errHandler {
		t.Errorf("handler returned empty list of digital currencys")
	}
}
