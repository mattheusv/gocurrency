package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msalcantara/gocurrency/currency"
)

//GetCurrencys get all traditional currencys
func GetCurrencys(w http.ResponseWriter, r *http.Request) {
	curr, err := currency.GetAllCurrency(currency.TRADITIONAL)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	currArray := map[string][]currency.Currency{
		"currencys": curr,
	}
	writeResponse(w, currArray)
}

//GetOneCurrency get traditional currency of parameter
func GetOneCurrency(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	curr, err := currency.GetCurrrency(param["currency"], currency.TRADITIONAL)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	currResponse := map[string]currency.Currency{
		"currency": curr,
	}
	writeResponse(w, currResponse)
}

//GetDigitalCurrencys get all digital currencys
func GetDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
	curr, err := currency.GetAllCurrency(currency.DIGITAL)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	currArray := map[string][]currency.Currency{
		"digital_currencys": curr,
	}
	writeResponse(w, currArray)
}

//GetOneDigitalCurrencys get digital currency of parameter
func GetOneDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	curr, err := currency.GetCurrrency(param["currency"], currency.DIGITAL)
	if err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}
	currResponse := map[string]currency.Currency{
		"digital_currency": curr,
	}
	writeResponse(w, currResponse)
}
