package serializers

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
)

type Serializer interface {
	Fill(e entities.Entity) error
	GetSerializerEntity() SerializerEntity
	Serialize(data interface{}, meta jsonapi.Meta) ([]byte, error)
}