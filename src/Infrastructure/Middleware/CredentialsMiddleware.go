package middleware

import (
	"net/http"

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
}

func (m *CredentialsMiddleware) CredentialsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userHeader := r.Header.Get(userHeader)
		platformHeader := r.Header.Get(platformHeader)

		if userHeader == "" || platformHeader == "" {
			// Missing required headers
			m.ErrorController.WriteHttpError(&domainerrors.UnauthorizeError{Err: "Unauthorized"}, w)
		} else {
			user := m.UserRepository.FindByID(userHeader)
			if user == nil {
				// User not found
			}

			platform := m.PlatformRepository.FindByID(platformHeader)
			if platform == nil {
				// Platform not found
			}

			if user.PlatformID != platform.ID {
				// User does not belongs to platform
			}
		}
	})
}