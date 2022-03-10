package didependencies

import (
	di "github.com/sarulabs/di/v2"

	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
)

var ControllersHttp = []di.Def{
	{
		Name:  "GetUserControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Get{
				Base: ctn.Get("BaseControllerHttp").(*controllerhttp.Base),
			}, nil
		},
	},
	{
		Name:  "BaseControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Base{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
	{
		Name:  "ErrorControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Error{
				Logger: ctn.Get("Logger").(logger.Logger),
				ErrorSerializer: ctn.Get("HttpErrorSerializer").(*serializers.ErrorSerializer),
			}, nil
		},
	},
}