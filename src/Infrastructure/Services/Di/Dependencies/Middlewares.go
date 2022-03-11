package didependencies

import (
	di "github.com/sarulabs/di/v2"

	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	middlewares "github.com/gnemes/go-users/Infrastructure/Middleware"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	repositories "github.com/gnemes/go-users/Domain/Model/Repositories"
)

var Middlewares = []di.Def{
	{
		Name:  "CredentialsMiddleware",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &middlewares.CredentialsMiddleware{
				Logger:             ctn.Get("Logger").(logger.Logger),
				UserRepository:     ctn.Get("UserRepository").(repositories.UserRepository),
				PlatformRepository: ctn.Get("PlatformRepository").(repositories.PlatformRepository),
				ErrorController:    ctn.Get("ErrorControllerHttp").(*controllerhttp.Error),
			}, nil
		},
	},
	{
		Name:  "JsonApiHeaderMiddleware",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &middlewares.JsonApiHeaderMiddleware{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
	{
		Name:  "TrimSlashMiddleware",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &middlewares.TrimSlashMiddleware{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
}