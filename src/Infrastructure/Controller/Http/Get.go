package controllerhttp

import (
	"net/http"

	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type Get struct {
	Logger logger.Logger
}

func (c *Get) Execute(w http.ResponseWriter, r *http.Request) {
	c.Logger.Debugf("Controller / Http / Get / Execute()")
	defer c.Logger.Debugf("Controller / Http / Get / Execute() ending...")
}