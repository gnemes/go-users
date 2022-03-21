package serializerentities

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

type UserProfile struct {
	*BaseSerializerEntity `json:"-"`
	
	Name     string  `json:"name"`
	LastName string  `json:"last-name"`
	Age      *int    `json:"age"`
	Phone    *string `json:"phone"`
}

func (se *UserProfile) GetName() string {
	return jsonapitypes.JsonapiUserProfileType
}

func (se *UserProfile) Fill(e entities.Entity) error {
	if userProfile, ok := e.(*entities.UserProfile); ok {
		se.Name     = userProfile.Name
		se.LastName = userProfile.LastName
		se.Age      = userProfile.Age
		se.Phone    = userProfile.Phone
		se.SetID(userProfile.ID)
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}