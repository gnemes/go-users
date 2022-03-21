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

	dummyPlatform := &entities.Platform{
		ID:   dummyPlatformID,
		Name: "Dummy Platform",
	}
		
	if userID == dummyID {
		dummyUser := &entities.User{
			ID:   dummyID,
			Username: "Dummy User",
			Platform: dummyPlatform,
		}
	
		return dummyUser
	}
	
	return nil
}