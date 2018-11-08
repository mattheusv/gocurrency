package errors

import "net/http"

type Error struct {
	Err string
}

func New(err string) (e Error) {
	return Error{Err: err}
}

func (e Error) Error() string {
	return e.Err
}

func (e Error) Write(w http.ResponseWriter, json []byte, httpStatus int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(json)
}
