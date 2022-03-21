package adminusecases

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	usecases "github.com/gnemes/go-users/Application/UseCases"
)

type GetUserUseCase struct {
	Logger     logger.Logger
	InputPort  *GetUserInputPort
	OutputPort *usecases.SingleOutputPort
}

func (uc *GetUserUseCase) Execute() error {
	uc.Logger.Debugf("Use Case / Admin / GetUserUseCase / Execute()")
	defer uc.Logger.Debugf("Use Case / Admin / GetUserUseCase / Build() Execute...")

	request := uc.InputPort.Request	

	// Dummy response
	platformResponse := &entities.Platform{
		ID: request.PlatformID,
		Name: "FooBarPlatform",
	}

	userResponse := &entities.User{
		ID: request.UserID,
		Username: "foo@bar.com",
		Platform: platformResponse,
	}

	uc.OutputPort.SetData(userResponse)

	return nil
}