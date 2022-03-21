package serializerentities

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

const (
	JsonapiPlatformName = "platform"
	JsonapiUserProfileName = "profile"
)

type User struct {
	*BaseSerializerEntity `json:"-"`
	
	Username   string `json:"username"`
}

func (se *User) GetName() string {
	return jsonapitypes.JsonapiUserType
}

func (se *User) Fill(e entities.Entity) error {
	if user, ok := e.(*entities.User); ok {
		se.Username   = user.Username
		se.SetID(user.ID)

		// Platform
		se.AddRelationship("PlatformSerializer", user.Platform, jsonapitypes.JsonapiPlatformType, JsonapiPlatformName)

		// User Profile
		if user.UserProfile != nil {
			se.AddRelationship("UserProfileSerializer", user.UserProfile, jsonapitypes.JsonapiUserProfileType, JsonapiUserProfileName)
		}
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}