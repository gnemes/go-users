package serializers

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
	serializerentities "github.com/gnemes/go-users/Infrastructure/Serializers/Entities"
)

type Serializer struct {
	Entity serializerentities.SerializerEntity
}

func (s *Serializer) Serialize(data interface{}, meta jsonapi.Meta) ([]byte, error) {
	if e, ok := data.(entities.Entity); ok {
		err := s.Entity.Fill(e)
		if err != nil {
			return nil, err
		}

		serializedData, errSerialize := s.Entity.SerializeEntity()
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
