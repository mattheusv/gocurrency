package api

import (
	"net/http"

	"github.com/urfave/negroni"
)

//commonMiddleware middleware to response allways application/json
func commonMiddleware() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Add("Content-Type", "application/json")
		next(w, r)
	})
}
