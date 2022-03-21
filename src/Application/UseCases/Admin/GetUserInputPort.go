package adminusecases

import (
	context "github.com/gnemes/go-users/Domain/Services/Context"
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type GetUserRequest struct {
	RequestID  string
	UserID     string
	PlatformID string
}

type GetUserInputPort struct {
	Logger  logger.Logger
	Context *context.Context

	Request *GetUserRequest
}

func (ip *GetUserInputPort) Build() error {
	ip.Logger.Debugf("Use Case / Admin / Input Port / GetUserInputPort / Build()")
	defer ip.Logger.Debugf("Use Case / Admin / Input Port / GetUserInputPort / Build() ending...")
	
	request := &GetUserRequest{
		RequestID: ip.Context.Get("RequestID").(string),
		UserID: ip.Context.Get("RequestedEntityID").(string),
		PlatformID: ip.Context.Get("Platform").(*entities.Platform).ID,
	}

	ip.Request = request

	return nil
}