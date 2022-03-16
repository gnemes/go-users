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
	request := uc.InputPort.Request	

	uc.Logger.Debugf("###### REQUEST :: %v", request)

	// Dummy response
	userResponse := &entities.User{
		ID: request.UserID,
		Username: "foo@bar.com",
		PlatformID: request.PlatformID,
	}

	uc.OutputPort.SetData(userResponse)

	return nil
}