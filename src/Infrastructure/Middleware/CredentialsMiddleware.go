package middleware

import (
	"net/http"

	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	di "github.com/gnemes/go-users/Infrastructure/Services/Di"
	domainerrors "github.com/gnemes/go-users/Domain/Errors"
)

const (
	UserHeader     = "X-USER-ID"
	PlatformHeader = "X-PLATFORM-NAME"
)

func CredentialsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get(UserHeader)
		platformName := r.Header.Get(PlatformHeader)

		if userID == "" || platformName == "" {
			// Missing required headers
			err := &domainerrors.UnauthorizeError{Err: "Unauthorized"}
			baseController := di.GetInstance().Get("BaseControllerHttp").(*controllerhttp.Base)
			baseController.WriteHttpError(err, w)
		} else {

		}
	})
}