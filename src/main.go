package main

import (
	di "github.com/gnemes/go-users/Infrastructure/Services/Di"

	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

func main() {
	logger := di.GetInstance().Get("Logger").(logger.Logger)
	logger.Infof("----- Starting Web Server... -----")
	defer logger.Infof("----- Ending Web Server -----")
	defer di.DeleteInstance()

	Listen()
}