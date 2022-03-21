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
	dummyUserProfileID := "102938"

	dummyPlatform := &entities.Platform{
		ID:   dummyPlatformID,
		Name: "Dummy Platform",
	}
		
	if userID == dummyID {
		age := 42
		phone := "1234566677"
		dummyUserProfile := &entities.UserProfile{
			ID:       dummyUserProfileID,
			Name:     "Dummy",
			LastName: "User",
			Age:      &age,
			Phone:    &phone,
		}

		dummyUser := &entities.User{
			ID:          dummyID,
			Username:    "DummyUser",
			Platform:    dummyPlatform,
			UserProfile: dummyUserProfile,
		}

		r.Logger.Debugf("USER REPOSITORY :: %v", dummyUser)
	
		return dummyUser
	}
	
	return nil
}