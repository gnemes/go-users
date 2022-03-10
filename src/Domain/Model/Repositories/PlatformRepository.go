package repositories

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
)

type PlatformRepository interface {
	FindByID(platformID string) *entities.Platform
}