package didependencies

import (
	"log"
	
	di "github.com/sarulabs/di/v2"

	// Config
	config "github.com/gnemes/go-users/Infrastructure/Services/Config"

	// Logger
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	loggerimpl "github.com/gnemes/go-users/Infrastructure/Services/Logger"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
	uuidImpl "github.com/gnemes/go-users/Infrastructure/Services/Uuid"
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
	{
		Name:  "Uuid",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			instance, errUuid := uuidImpl.NewUuid()
			if errUuid != nil {
				log.Fatalf("Could not instantiate uuid %s", errUuid.Error())
			}
			return instance.(uuid.Uuid), nil
		},
	},
}
