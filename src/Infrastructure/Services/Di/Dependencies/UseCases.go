package didependencies

import (
	di "github.com/sarulabs/di/v2"

	adminusecases "github.com/gnemes/go-users/Application/UseCases/Admin"
	context "github.com/gnemes/go-users/Domain/Services/Context"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	usecases "github.com/gnemes/go-users/Application/UseCases"
)

var UseCases = []di.Def{
	{
		Name:  "AdminGetUserInputPort",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &adminusecases.GetUserInputPort{
				Logger:  ctn.Get("Logger").(logger.Logger),
				Context: ctn.Get("Context").(*context.Context),
			}, nil
		},
	},
	{
		Name:  "AdminGetUserOutputPort",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &usecases.SingleOutputPort{}, nil
		},
	},
	{
		Name:  "AdminGetUserUseCase",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &adminusecases.GetUserUseCase{
				Logger:     ctn.Get("Logger").(logger.Logger),
				InputPort:  ctn.Get("AdminGetUserInputPort").(*adminusecases.GetUserInputPort),
				OutputPort: ctn.Get("AdminGetUserOutputPort").(*usecases.SingleOutputPort),
			}, nil
		},
	},
}