package api

import (
	"errors"
	"net/http"
)

//notFound default error handler to 404 page not found
func notFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeError(w, errors.New("404 endpoint not found"), http.StatusNotFound)
	})
}
