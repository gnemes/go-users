package main

import (
	di "github.com/gnemes/go-users/Infrastructure/Services/Di"

	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

func main() {
	container := di.BuildDi()

	logger := container.Get("Logger").(logger.Logger)
	logger.Infof("----- Starting Web Server... -----")
	defer logger.Infof("----- Ending Web Server -----")
	defer container.Delete()

	listen(container)
}