package middleware

import (
	"net/http"

	"github.com/gorilla/mux"

	context "github.com/gnemes/go-users/Domain/Services/Context"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type RequestedEntityIDMiddleware struct {
	Logger  logger.Logger
	Context *context.Context
}

func (m *RequestedEntityIDMiddleware) Execute(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Debugf("Middleware / RequestedEntityIDMiddleware()")
		defer m.Logger.Debugf("Middleware / RequestedEntityIDMiddleware() ending...")
		
		if mux.Vars(r)["id"] != "" {
			entityID := mux.Vars(r)["id"]
			m.Context.Add("RequestedEntityID", entityID)
		}

		next(w, r)
	}
}