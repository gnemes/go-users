package di

import (
	"github.com/sarulabs/di"

	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

var ControllersHttp = []di.Def{
	{
		Name:  "GetUserControllerHttp",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Get{
				Base: ctn.Get("BaseControllerHttp").(*controllerhttp.Base),
			}, nil
		},
	},
	{
		Name:  "BaseControllerHttp",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Base{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
}