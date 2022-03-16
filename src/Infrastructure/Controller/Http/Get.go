package controllerhttp

import (
	"net/http"

	usecases "github.com/gnemes/go-users/Domain/UseCases"
)

type Get struct {
	*Base
	AdminInputPort   usecases.InputPort
	AdminUseCase     usecases.UseCase
	Presenter        usecases.Presenter
}

func (c *Get) Execute(w http.ResponseWriter, r *http.Request) {
	c.Logger.Debugf("Controller / Http / Get / Execute() request: %s", c.Context.Get("RequestID").(string))
	defer c.Logger.Debugf("Controller / Http / Get / Execute() request: %s ending...", c.Context.Get("RequestID").(string))

	var currentInputPort usecases.InputPort
	var currentUseCase usecases.UseCase

	// Change this to check if current context if from an admin user or not
	currentInputPort = c.AdminInputPort
	currentUseCase = c.AdminUseCase

	// Build input port
	errInputPort := currentInputPort.Build()
	if errInputPort == nil {
		// Execute use case
		errUseCase := currentUseCase.Execute()
		if errUseCase == nil {
			// Present response
			response, errPresenter := c.Presenter.Present()
			if errPresenter == nil {
				// Write response
				w.Write(response)
			} else {
				c.Logger.Errorf("ERROR: Error presenting response: %s", errPresenter.Error())
				c.ErrorController.WriteHttpError(errPresenter, w)		
			}
		} else {
			c.Logger.Errorf("ERROR: Error on use case: %s", errUseCase.Error())
			c.ErrorController.WriteHttpError(errUseCase, w)	
		}
	} else {
		c.Logger.Errorf("ERROR: Error building input port: %s", errInputPort.Error())
		c.ErrorController.WriteHttpError(errInputPort, w)
	}
}