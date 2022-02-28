package config

import (
	"os"
	"strconv"
	"sync"
)

const TESTING_ENV = "test"

var once sync.Once

// Config is struct for application settings
type Config struct {
	Environment string

	MySQLUser     string
	MySQLPassword string
	MySQLHost     string
	MySQLPort     int
	MySQLDatabase string

	LoggerConsoleEnable   bool
	LoggerConsoleLevel    string
	LoggerFileEnable      bool
	LoggerFileLevel       string
	LoggerFileLocation    string

	// RabbitConnection    string
	// RabbitUserExchange  string
	// RabbitRecoExchange  string
	// RabbitRulesExchange string
	// RabbitQueue         string
	// RabbitConsumer      string
}

// Global variable
var instance *Config

// IsTestingEnvironment Flag to check if it is a testing execution
func IsTestingEnvironment() bool {
	return instance.Environment == TESTING_ENV
}

// Get get config from environment variables
func Get() *Config {
	once.Do(func() {
		instance = &Config{
			Environment: getenvStr("ENVIRONMENT"),

			MySQLUser:     getenvStr("MYSQL_USERNAME"),
			MySQLPassword: getenvStr("MYSQL_PASSWORD"),
			MySQLHost:     getenvStr("MYSQL_HOST"),
			MySQLPort:     getenvInt("MYSQL_PORT"),
			MySQLDatabase: getenvStr("MYSQL_DATABASE"),

			LoggerConsoleEnable: getenvStr("LOGGER_CONSOLE_ENABLE") == "true",
			LoggerConsoleLevel:  getenvStr("LOGGER_CONSOLE_LEVEL"),
			LoggerFileEnable:    getenvStr("LOGGER_FILE_ENABLE") == "true",
			LoggerFileLevel:     getenvStr("LOGGER_FILE_LEVEL"),
			LoggerFileLocation:  getenvStr("LOGGER_FILE_LOCATION"),

			// RabbitConnection:    getenvStr("RABBIT_CONNECTION"),
			// RabbitUserExchange:  getenvStr("RABBIT_USERS_EXCHANGE"),
			// RabbitRecoExchange:  getenvStr("RABBIT_RECO_EXCHANGE"),
			// RabbitRulesExchange: getenvStr("RABBIT_RULES_EXCHANGE"),
			// RabbitQueue:         getenvStr("RABBIT_QUEUE"),
			// RabbitConsumer:      getenvStr("RABBIT_CONSUMER"),
		}
	})

	return instance
}

// getenvStr Get env variable as string
func getenvStr(key string) string {
	v := os.Getenv(key)
	return v
}

// getenvInt Get env variable as int
func getenvInt(key string) int {
	s := getenvStr(key)
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// getenvBool Get env variable as bool
func getenvBool(key string) bool {
	s := getenvStr(key)
	if s == "" {
		return false
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return v
}