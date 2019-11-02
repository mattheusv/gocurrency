package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/msalcantara/gocurrency/currency"
	"github.com/urfave/negroni"
)

type API struct {
	c *currency.CurrencyImp
}

type Server struct {
	api  API
	port int
}

func NewServer(port int) *Server {
	return &Server{port: port, api: API{
		c: currency.NewCurrency(),
	}}
}

type response struct {
	Data       interface{} `json:"data,omitempty"`
	Error      error       `json:"error,omitempty"`
	StatusCode int         `json:"-"`
}

type apiFunc func(r *http.Request) response

func (server *Server) Run() error {
	//init router
	r := mux.NewRouter().StrictSlash(true)
	server.api.Register(r)

	//init middleware
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(commonMiddleware())

	n.UseHandler(r)

	log.Printf("server running on: %d\n", server.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", server.port), n); err != nil {
		return err
	}
	return nil
}

func (api *API) Register(r *mux.Router) {
	wrap := func(f apiFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			response := f(r)
			if response.Error != nil {
				log.Printf("ERROR: %v\n", response.Error)
				if response.StatusCode == 0 {
					response.StatusCode = http.StatusInternalServerError
				}
				b, err := json.Marshal(&response)
				if err != nil {
					log.Printf("ERROR: error to marshal response: %v\n", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(b)
				return
			}
			if response.Data != nil {
				if response.StatusCode == 0 {
					response.StatusCode = http.StatusOK
				}
				b, err := json.Marshal(&response)
				if err != nil {
					log.Printf("ERROR: error to marshal response: %v\n", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(b)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
	r.HandleFunc("/api/currencys", wrap(api.currencys)).Methods("GET")
	r.HandleFunc("/api/currencys/{currency}", wrap(api.currency)).Methods("GET")
	r.HandleFunc("/api/digital/currencys", wrap(api.digitalCurrencys)).Methods("GET")
	r.HandleFunc("/api/digital/currencys/{currency}", wrap(api.digitalCurrency)).Methods("GET")

	// 4xx errors routers
	r.NotFoundHandler = wrap(api.notFound)
}

func (api *API) currencys(r *http.Request) response {
	c, err := api.c.Traditional.GetAll()
	if err != nil {
		return response{
			Error:      err,
			StatusCode: http.StatusBadRequest,
		}
	}
	data := map[string][]currency.Currency{
		"currencys": c,
	}
	return response{Data: data}
}

func (api *API) currency(r *http.Request) response {
	param := mux.Vars(r)
	c, err := api.c.Traditional.GetOne(param["currency"])
	if err != nil {
		return response{Error: err, StatusCode: http.StatusBadRequest}
	}
	data := map[string]currency.Currency{
		"currency": c,
	}
	return response{Data: data}
}

func (api *API) digitalCurrencys(r *http.Request) response {
	c, err := api.c.Digital.GetAll()
	if err != nil {
		return response{Error: err, StatusCode: http.StatusBadRequest}
	}
	data := map[string][]currency.Currency{
		"digital_currencys": c,
	}
	return response{Data: data}
}

func (api *API) digitalCurrency(r *http.Request) response {
	param := mux.Vars(r)
	c, err := api.c.Digital.GetOne(param["currency"])
	if err != nil {
		return response{Error: err, StatusCode: http.StatusBadRequest}
	}
	data := map[string]currency.Currency{
		"digital_currency": c,
	}
	return response{Data: data}
}

func (api *API) notFound(r *http.Request) response {
	return response{Error: errors.New("404 endpoint not found"), StatusCode: http.StatusNotFound}
}

func commonMiddleware() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Add("Content-Type", "application/json")
		next(w, r)
	})
}
