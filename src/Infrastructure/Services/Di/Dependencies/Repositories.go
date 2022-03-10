package didependencies

import (
	di "github.com/sarulabs/di/v2"

	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	repositoriesimpl "github.com/gnemes/go-users/Infrastructure/Model/Repositories"
)

var Repositories = []di.Def{
	{
		Name:  "UserRepository",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &repositoriesimpl.DummyUserRepository{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
	{
		Name:  "PlatformRepository",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &repositoriesimpl.DummyPlatformRepository{
				Logger: ctn.Get("Logger").(logger.Logger),
			}, nil
		},
	},
}