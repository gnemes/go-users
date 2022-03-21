package serializerentities

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

type Platform struct {
	*BaseSerializerEntity `json:"-"`
	
	Name string `json:"name"`
}

func (se *Platform) GetName() string {
	return jsonapitypes.JsonapiPlatformType
}

func (se *Platform) Fill(e entities.Entity) error {
	if platform, ok := e.(*entities.Platform); ok {
		se.Name   = platform.Name
		se.SetID(platform.ID)
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}