package controllerhttp

import (
	"net/http"
)

type Get struct {
	*Base
}

func (c *Get) Execute(w http.ResponseWriter, r *http.Request) {
	c.Logger.Debugf("Controller / Http / Get / Execute() request: %s", c.Context.Get("RequestID").(string))
	defer c.Logger.Debugf("Controller / Http / Get / Execute() request: %s ending...", c.Context.Get("RequestID").(string))
}