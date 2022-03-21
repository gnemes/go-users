package serializers

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
	serializersdomain "github.com/gnemes/go-users/Domain/Serializers"
)

type Serializer struct {
	Entity serializersdomain.SerializerEntity
}

func (s *Serializer) Serialize(data interface{}, meta jsonapi.Meta) ([]byte, error) {
	if e, ok := data.(entities.Entity); ok {
		err := s.Entity.Fill(e)
		if err != nil {
			return nil, err
		}

		serializedData, errSerialize := s.Entity.Serialize()
		if errSerialize != nil {
			return nil, errSerialize
		}
	
		if serializedData != nil {
			return serializedData, nil
		} else {
			return jsonapi.Marshal(s.Entity, meta)
		}
	} else {
		return nil, errors.New("Invalid data to serialize")
	}
}

func (s *Serializer) Fill(e entities.Entity) error {
	return s.Entity.Fill(e)
}

func (s *Serializer) GetSerializerEntity() serializersdomain.SerializerEntity {
	return s.Entity
}