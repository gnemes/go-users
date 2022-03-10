package repositories

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type DummyUserRepository struct {
	Logger logger.Logger
}

func (r *DummyUserRepository) FindByID(userID string) *entities.User {
	r.Logger.Debugf("Repository / User / Dummy / FindByID()")
	defer r.Logger.Debugf("Repository / User / Dummy / FindByID ending...")

	dummyID := "5678"
	dummyPlatformID := "1234"

	if userID == dummyID {
		dummyUser := &entities.User{
			ID:   dummyID,
			PlatformID: dummyPlatformID,
			Username: "Dummy User",
		}
	
		return dummyUser
	}
	
	return nil
}