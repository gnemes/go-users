package controllerhttp

import (
	"net/http"

	domainerrors "github.com/gnemes/go-users/Domain/Errors"
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
)

type Error struct {
	Logger          logger.Logger
	Uuid            uuid.Uuid
	ErrorSerializer *serializers.Serializer
}

// WriteError in response writer
func (c *Error) WriteHttpError(err error, w http.ResponseWriter) {
	var httpError entities.Error

	switch err.(type) {
	case *domainerrors.NotImplementedError:
		httpError = entities.Error{
			c.Uuid.New(),
			501,
			"NOT_IMPLEMENTED",
			"Not implemented",
			err.Error(),
		}
	case *domainerrors.DomainError:
		httpError = entities.Error{
			c.Uuid.New(),
			500,
			"DOMAIN_ERROR",
			"Domain Error",
			err.Error(),
		}
	case *domainerrors.NotFoundError:
		httpError = entities.Error{
			c.Uuid.New(),
			404,
			"NOT FOUND",
			"Not Found",
			err.Error(),
		}
	case *domainerrors.ForbiddenError:
		httpError = entities.Error{
			c.Uuid.New(),
			403,
			"FORBIDDEN",
			"Forbidden",
			err.Error(),
		}
	case *domainerrors.UnauthorizeError:
		httpError = entities.Error{
			c.Uuid.New(),
			403,
			"UNAUTHORIZE",
			"Unauthorize",
			err.Error(),
		}
	case *domainerrors.MissingFilterError:
		httpError = entities.Error{
			c.Uuid.New(),
			400,
			"MISSING_FILTER",
			"Missing Filter",
			err.Error(),
		}
	case *domainerrors.BadRequestError:
		httpError = entities.Error{
			c.Uuid.New(),
			400,
			"BAD_REQUEST",
			"Bad Request",
			(*err.(*domainerrors.BadRequestError)).Message,
		}
	default:
		errStr := "Something went wrong: " + err.Error()

		httpError = entities.Error{
			c.Uuid.New(),
			500,
			"INTERNAL_SERVER_ERROR",
			"Internal Server Error",
			errStr,
		}
	}

	jsonresult, _ := c.ErrorSerializer.Serialize(&httpError, nil)
	w.WriteHeader(httpError.Status)
	w.Write(jsonresult)
}