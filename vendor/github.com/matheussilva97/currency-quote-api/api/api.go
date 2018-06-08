package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/matheussilva97/currency-quote-api/scraping"

	"github.com/gorilla/mux"
)

// GetCurrencys get all currencys
func GetCurrencys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	curr := []scraping.Currency{}
	done := make(chan bool)
	go scraping.GetAllCurrency(&curr, 0, done)
	<-done
	currArray := map[string][]scraping.Currency{
		"currencys": curr,
	}
	json.NewEncoder(w).Encode(currArray)
}

// GetOneCurrency get currency of parameter
func GetOneCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	curr := scraping.GetCurrrency(param["currency"], 0)
	currResponse := map[string]scraping.Currency{
		"currency": curr,
	}
	json.NewEncoder(w).Encode(currResponse)
}

// GetDigitalCurrencys get all digital currencys
func GetDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contenty-Type", "application/json")
	curr := []scraping.Currency{}
	done := make(chan bool)
	go scraping.GetAllCurrency(&curr, 1, done)
	<-done
	currArray := map[string][]scraping.Currency{
		"digital_currencys": curr,
	}
	json.NewEncoder(w).Encode(currArray)
}

// GetOneDigitalCurrencys get digital currency of parameter
func GetOneDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	curr := scraping.GetCurrrency(param["currency"], 1)
	currResponse := map[string]scraping.Currency{
		"digital_currency": curr,
	}
	json.NewEncoder(w).Encode(currResponse)
}

// Routers provide endpoints
func Routers() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/currencys", GetCurrencys).Methods("GET")
	router.HandleFunc("/api/currencys/{currency}", GetOneCurrency).Methods("GET")
	router.HandleFunc("/api/digital/currencys", GetDigitalCurrencys).Methods("GET")
	router.HandleFunc("/api/digital/currencys/{currency}", GetOneDigitalCurrencys).Methods("GET")

	fmt.Println("Server running on http://127.0.0.1:8080/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
