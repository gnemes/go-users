package middleware

import (
    "net/http"

    context "github.com/gnemes/go-users/Domain/Services/Context"
    logger "github.com/gnemes/go-users/Domain/Services/Logger"
    uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
)

type RequestIDMiddleware struct {
    Logger  logger.Logger
    Context *context.Context
    Uuid    uuid.Uuid
}

func (m *RequestIDMiddleware) Execute(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        m.Logger.Debugf("Middleware / RequestIDMiddleware()")
		defer m.Logger.Debugf("Middleware / RequestIDMiddleware() ending...")

        requestID := m.Uuid.New()
        m.Context.Add("RequestID", requestID)

        m.Logger.Debugf("Request ID: %s", requestID)
        next(w, r)
    }
}
