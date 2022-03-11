package middleware

import (
    "strings"
    "net/http"

    logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type TrimSlashMiddleware struct {
    Logger logger.Logger
}

func (m *TrimSlashMiddleware) Execute(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        m.Logger.Debugf("Middleware / TrimSlashMiddleware()")
		defer m.Logger.Debugf("Middleware / TrimSlashMiddleware() ending...")

        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
        next.ServeHTTP(w, r)
    })
}
