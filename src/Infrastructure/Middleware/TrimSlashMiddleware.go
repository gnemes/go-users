package middleware

import (
    "strings"
    "net/http"
)

func TrimSlashMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
        next.ServeHTTP(w, r)
    })
}
