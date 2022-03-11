package controllerhttp

import (
	"net/http"

	"time"
)

type Get struct {
	*Base
}

func (c *Get) Execute(w http.ResponseWriter, r *http.Request) {
	c.Logger.Debugf("Controller / Http / Get / Execute() request: %s", c.Context.Get("RequestID").(string))
	defer c.Logger.Debugf("Controller / Http / Get / Execute() request: %s ending...", c.Context.Get("RequestID").(string))

	time.Sleep(8 * time.Second)
}