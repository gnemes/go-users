package serializerentities

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

type User struct {
	*BaseSerializerEntity `json:"-"`
	
	Username   string `json:"username"`
	PlatformID string `json:"platform-id"`
}

func (se *User) GetName() string {
	return jsonapitypes.JsonapiUserType
}

func (se *User) Fill(e entities.Entity) error {
	if user, ok := e.(*entities.User); ok {
		se.Username   = user.Username
		se.PlatformID = user.PlatformID
		se.SetID(user.ID)
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}