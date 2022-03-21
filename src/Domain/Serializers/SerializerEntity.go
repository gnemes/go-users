package serializers

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
)

type SerializerEntity interface {
	GetID() string
	SetID(id string) error
	Fill(e entities.Entity) error
	Serialize() ([]byte, error)
}