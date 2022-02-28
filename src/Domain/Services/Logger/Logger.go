package logger

// LoggerConfiguration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type LoggerConfiguration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
}

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}