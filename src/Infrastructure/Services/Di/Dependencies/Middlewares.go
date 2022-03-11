package didependencies

import (
	di "github.com/sarulabs/di/v2"

	context "github.com/gnemes/go-users/Domain/Services/Context"
	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	middlewares "github.com/gnemes/go-users/Infrastructure/Middleware"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	repositories "github.com/gnemes/go-users/Domain/Model/Repositories"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
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
				Context:            ctn.Get("Context").(*context.Context),
			}, nil
		},
	},
	{
		Name:  "RequestQueryParserMiddleware",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &middlewares.RequestQueryParserMiddleware{
				Logger:             ctn.Get("Logger").(logger.Logger),
				ErrorController:    ctn.Get("ErrorControllerHttp").(*controllerhttp.Error),
				Context:            ctn.Get("Context").(*context.Context),
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
	{
		Name:  "RequestIDMiddleware",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &middlewares.RequestIDMiddleware{
				Logger:  ctn.Get("Logger").(logger.Logger),
				Context: ctn.Get("Context").(*context.Context),
				Uuid:    ctn.Get("Uuid").(uuid.Uuid),
			}, nil
		},
	},
}