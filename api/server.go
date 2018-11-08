package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msalcantara/currency-quote-api/config"
	"github.com/urfave/negroni"
)

//StartServer start server
func StartServer() {
	r := setHandler()
	http.Handle("/", r)
	addr := fmt.Sprintf("%s:%d", config.ReportConfig.HTTP.Host, config.ReportConfig.HTTP.Port)
	log.SetPrefix("[report] ")
	log.Printf("Server Running on https://%s/ (Press CTRL+C to quit)", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error to start server %v\n", err)
	}
}

func setHandler() http.Handler {
	r := initRouter()
	n := initNegroni()
	r.HandleFunc("/api/currencys", GetCurrencys).Methods("GET")
	r.HandleFunc("/api/currencys/{currency}", GetOneCurrency).Methods("GET")
	r.HandleFunc("/api/digital/currencys", GetDigitalCurrencys).Methods("GET")
	r.HandleFunc("/api/digital/currencys/{currency}", GetOneDigitalCurrencys).Methods("GET")
	n.UseHandler(r)
	return n
}

func initRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.NotFoundHandler = notFound()
	return r
}

func initNegroni() *negroni.Negroni {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(commonMiddleware())
	return n
}
