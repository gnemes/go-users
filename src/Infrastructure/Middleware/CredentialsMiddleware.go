package middleware

import (
	"net/http"

	context "github.com/gnemes/go-users/Domain/Services/Context"
	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	domainerrors "github.com/gnemes/go-users/Domain/Errors"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	repositories "github.com/gnemes/go-users/Domain/Model/Repositories"
)

const (
	userHeader     = "X-USER-ID"
	platformHeader = "X-PLATFORM-ID"
)

type CredentialsMiddleware struct {
	Logger             logger.Logger
	UserRepository     repositories.UserRepository
	PlatformRepository repositories.PlatformRepository
	ErrorController    *controllerhttp.Error
	Context            *context.Context
}

func (m *CredentialsMiddleware) Execute(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Debugf("Middleware / CredentialsMiddleware()")
		defer m.Logger.Debugf("Middleware / CredentialsMiddleware() ending...")

		userHeader := r.Header.Get(userHeader)
		platformHeader := r.Header.Get(platformHeader)

		if userHeader == "" || platformHeader == "" {
			// Missing required headers
			m.Logger.Errorf("Unauthorized request. Missing headers.")
			m.ErrorController.WriteHttpError(&domainerrors.UnauthorizeError{Err: "Unauthorized"}, w)
		} else {
			user := m.UserRepository.FindByID(userHeader)
			if user == nil {
				// User not found
				m.Logger.Errorf("Unauthorized request. User not found.")
				m.ErrorController.WriteHttpError(&domainerrors.UnauthorizeError{Err: "Unauthorized"}, w)
				return
			}
			
			platform := m.PlatformRepository.FindByID(platformHeader)
			if platform == nil {
				// Platform not found
				m.Logger.Errorf("Unauthorized request. Platform not found.")
				m.ErrorController.WriteHttpError(&domainerrors.UnauthorizeError{Err: "Unauthorized"}, w)
				return
			}

			if user.Platform.ID != platform.ID {
				// User does not belongs to platform
				m.Logger.Errorf("Unauthorized request. User does not belongs to platform.")
				m.ErrorController.WriteHttpError(&domainerrors.UnauthorizeError{Err: "Unauthorized"}, w)
				return
			}
			
			m.Context.Add("User", user)
			m.Context.Add("Platform", platform)

			next(w, r)
		}
	}
}