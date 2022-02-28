package logger

import (
	"errors"
	
	domainlogger "github.com/gnemes/go-users/Domain/Services/Logger"
)

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
)

var (
	errInvalidLoggerInstance = errors.New("Invalid logger instance")
)

//NewLogger returns an instance of logger
func NewLogger(
	config domainlogger.LoggerConfiguration,
	loggerInstance int,
) (domainlogger.Logger, error) {
	switch loggerInstance {
	case InstanceZapLogger:
		logger, err := newZapLogger(config)
		if err != nil {
			return nil, err
		}
		return logger, nil
	default:
		return nil, errInvalidLoggerInstance
	}
}