package repositories

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type DummyPlatformRepository struct {
	Logger logger.Logger
}

func (r *DummyPlatformRepository) FindByID(platformID string) *entities.Platform {
	r.Logger.Debugf("Repository / Platform / Dummy / FindByID()")
	defer r.Logger.Debugf("Repository / Platform / Dummy / FindByID ending...")

	dummyID := "1234"

	if platformID == dummyID {
		dummyPlatform := &entities.Platform{
			ID:   dummyID,
			Name: "Dummy Platform",
		}
	
		return dummyPlatform
	}
	
	return nil
}