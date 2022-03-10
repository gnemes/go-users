package repositories

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
)

type UserRepository interface {
	FindByID(userID string) *entities.User
}