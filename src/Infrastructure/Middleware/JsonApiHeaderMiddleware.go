package middleware

import (
	"net/http"
)

func ApplicationJsonMiddleware(next http.Handler) http.Handler {
	const contentType = "application/json"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)
		next.ServeHTTP(w, r)
	})
}
