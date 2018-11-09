package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/msalcantara/gocurrency/errors"
)

type response struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func (r response) Bytes() []byte {
	b, _ := json.Marshal(&r)
	return b
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	if e, ok := err.(errors.Error); ok {
		r := &response{
			Data:  map[string]string{},
			Error: e.Error(),
		}
		w.WriteHeader(statusCode)
		w.Write(r.Bytes())
		return
	}
	r := &response{
		Data:  map[string]string{},
		Error: "internal error ocurred",
	}
	log.Printf("error: %v\n", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(r.Bytes())
}

func writeResponse(w http.ResponseWriter, data interface{}) {
	r := response{
		Data:  data,
		Error: "",
	}
	w.Write(r.Bytes())
}
