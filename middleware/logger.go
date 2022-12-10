package middleware

import (
	"log"
	"net/http"
)

func LogRequestWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("request: " + r.URL.String())
		h.ServeHTTP(w, r)
	})
}
