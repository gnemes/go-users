package controllerhttp

import (
	context "github.com/gnemes/go-users/Domain/Services/Context"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type Base struct {
	Logger          logger.Logger
	Context         *context.Context
	ErrorController *Error
}