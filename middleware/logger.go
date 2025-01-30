package middleware

import (
	"log"
	"net/http"
)

func LogRequestWrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value(ctxKeyUsername)
		if username == nil {
			username = "unknown"
		}
		log.Printf("Username %s requested: %s", username, r.URL.String())
		h.ServeHTTP(w, r)
	}
}
