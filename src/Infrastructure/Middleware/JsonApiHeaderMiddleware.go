package middleware

import (
	"net/http"

	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type JsonApiHeaderMiddleware struct {
	Logger logger.Logger
}

const (
	contentType = "application/json"
)

func (m *JsonApiHeaderMiddleware) JsonApiHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)
		next.ServeHTTP(w, r)
	})
}
