package controllerhttp

import (
	"net/http"

	domainerrors "github.com/gnemes/go-users/Domain/Errors"
	httperror "github.com/gnemes/go-users/Infrastructure/Controller/Http/Error"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
)

type Error struct {
	Logger          logger.Logger
	ErrorSerializer *serializers.ErrorSerializer
}

// WriteError in response writer
func (c *Error) WriteHttpError(err error, w http.ResponseWriter) {
	var httpError httperror.HttpError

	switch err.(type) {
	case *domainerrors.NotImplementedError:
		httpError = httperror.HttpError{
			501,
			"NOT_IMPLEMENTED",
			"Not implemented",
			err.Error(),
		}
	case *domainerrors.DomainError:
		httpError = httperror.HttpError{
			500,
			"DOMAIN_ERROR",
			"Domain Error",
			err.Error(),
		}
	case *domainerrors.NotFoundError:
		httpError = httperror.HttpError{
			404,
			"NOT FOUND",
			"Not Found",
			err.Error(),
		}
	case *domainerrors.ForbiddenError:
		httpError = httperror.HttpError{
			403,
			"FORBIDDEN",
			"Forbidden",
			err.Error(),
		}
	case *domainerrors.UnauthorizeError:
		httpError = httperror.HttpError{
			403,
			"UNAUTHORIZE",
			"Unauthorize",
			err.Error(),
		}
	case *domainerrors.MissingFilterError:
		httpError = httperror.HttpError{
			400,
			"MISSING_FILTER",
			"Missing Filter",
			err.Error(),
		}
	case *domainerrors.BadRequestError:
		httpError = httperror.HttpError{
			400,
			"BAD_REQUEST",
			"Bad Request",
			(*err.(*domainerrors.BadRequestError)).Message,
		}
	default:
		errStr := "Something went wrong: " + err.Error()

		httpError = httperror.HttpError{
			500,
			"INTERNAL_SERVER_ERROR",
			"Internal Server Error",
			errStr,
		}
	}

	jsonresult, _ := c.ErrorSerializer.Serialize(httpError)
	w.WriteHeader(httpError.Status)
	w.Write(jsonresult)
}