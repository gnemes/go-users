package di

import (
	"log"
	
	"github.com/sarulabs/di"

	// Config
	config "github.com/gnemes/go-users/Infrastructure/Services/Config"

	// Logger
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	loggerimpl "github.com/gnemes/go-users/Infrastructure/Services/Logger"
)

var Base = []di.Def{
	{
		Name:  "Logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			// Logger configuration
			loggerConfig := logger.LoggerConfiguration{
				EnableConsole:     config.Get().LoggerConsoleEnable,
				ConsoleLevel:      config.Get().LoggerConsoleLevel,
				ConsoleJSONFormat: true,
				EnableFile:        config.Get().LoggerFileEnable,
				FileLevel:         config.Get().LoggerFileLevel,
				FileJSONFormat:    true,
				FileLocation:      config.Get().LoggerFileLocation,
			}

			// Create logger instance
			l, err := loggerimpl.NewLogger(loggerConfig, loggerimpl.InstanceZapLogger)
			if err != nil {
				log.Fatalf("Could not instantiate log %s", err.Error())
			}
			return l.(logger.Logger), nil
		},
	},
}
