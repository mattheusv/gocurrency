package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/msalcantara/currency-quote-api/config"
	"github.com/urfave/negroni"
)

//StartServer start server
func StartServer() {
	r := setHandler()
	http.Handle("/", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(config.ReportConfig.HTTP.Port)
	}
	addr := fmt.Sprintf("%s:%s", config.ReportConfig.HTTP.Host, port)
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
