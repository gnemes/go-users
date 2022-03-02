package httpclientutils

import (
	"net/http"

	domainerrors "github.com/gnemes/go-users/Domain/Errors"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
)

// WriteError in response writer
func WriteError(err error, w http.ResponseWriter) {
	errorserializer := &serializer.ErrorSerializer{}

	switch e := err.(type) {
	case *domainerrors.NotImplementedError:
		jsonresult, _ := errorserializer.Serialize(
			"501",
			"NOT_IMPLEMENTED",
			"Not implemented",
			err.Error(),
		)
		w.WriteHeader(501)
		w.Write(jsonresult)
	case *domainerrors.DomainError:
		jsonresult, _ := errorserializer.Serialize(
			"500",
			"DOMAIN_ERROR",
			"Domain Error",
			err.Error(),
		)
		w.WriteHeader(500)
		w.Write(jsonresult)
	case *domainerrors.InvalidCsvFileError:
		invalidCsvSerializer := &serializer.InvalidCsvErrorSerializer{}
		jsonresult, _ := invalidCsvSerializer.Serialize(
			"400",
			"DOMAIN_ERROR",
			"Invalid csv file",
			"Invalid csv file.",
			err.(*domainerrors.InvalidCsvFileError).ErrorFileUrl,
		)
		w.WriteHeader(400)
		w.Write(jsonresult)
	case *domainerrors.NotFoundError:
		jsonresult, _ := errorserializer.Serialize(
			"404",
			"NOT FOUND",
			"Not Found",
			err.Error(),
		)
		w.WriteHeader(404)
		w.Write(jsonresult)
	case *domainerrors.ForbiddenError:
		jsonresult, _ := errorserializer.Serialize(
			"403",
			"FORBIDDEN",
			"Forbidden",
			err.Error(),
		)
		w.WriteHeader(403)
		w.Write(jsonresult)
	case *domainerrors.UnauthorizeError:
		jsonresult, _ := errorserializer.Serialize(
			"403",
			"UNAUTHORIZE",
			"Unauthorize",
			err.Error(),
		)
		w.WriteHeader(403)
		w.Write(jsonresult)
	case *domainerrors.MissingFilterError:
		jsonresult, _ := errorserializer.Serialize(
			"400",
			"MISSING_FILTER",
			"Missing Filter",
			err.Error(),
		)
		w.WriteHeader(400)
		w.Write(jsonresult)
	case *domainerrors.BadRequestError:
		jsonresult, _ := errorserializer.Serialize(
			"400",
			"BAD_REQUEST",
			"Bad Request",
			(*err.(*domainerrors.BadRequestError)).Message,
		)
		w.WriteHeader(400)
		w.Write(jsonresult)
	default:
		errStr := "Something went wrong: " + err.Error()

		jsonresult, _ := errorserializer.Serialize(
			"500",
			"INTERNAL_SERVER_ERROR",
			"Internal Server Error",
			errStr,
		)
		w.WriteHeader(500)
		w.Write(jsonresult)
	}
}
