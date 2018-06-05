package api

import (
	"currency-quote-api/scraping"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCurrencys get all
func GetCurrencys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	curr := []scraping.Currency{}
	done := make(chan bool)
	go scraping.GetAllCurrency(&curr, done)
	<-done
}

func GetOneCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	scraping.GetCurrrency(param["currency"])
}
