package api

import (
	"currency-quote-api/scraping"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCurrencys get all currencys
func GetCurrencys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	curr := []scraping.Currency{}
	done := make(chan bool)
	go scraping.GetAllCurrency(&curr, done)
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
	curr := scraping.GetCurrrency(param["currency"])
	json.NewEncoder(w).Encode(curr)
}

// Routers provide endpoints
func Routers() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/currencys", GetCurrencys).Methods("GET")
	router.HandleFunc("/api/currencys/{currency}", GetOneCurrency).Methods("GET")

	fmt.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
