package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msalcantara/gocurrency/currency"
)

//getCurrencys get all traditional currencys
func getCurrencys(w http.ResponseWriter, r *http.Request) {
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

//getOneCurrency get traditional currency of parameter
func getOneCurrency(w http.ResponseWriter, r *http.Request) {
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

//getDigitalCurrencys get all digital currencys
func getDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
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

//getOneDigitalCurrencys get digital currency of parameter
func getOneDigitalCurrencys(w http.ResponseWriter, r *http.Request) {
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
