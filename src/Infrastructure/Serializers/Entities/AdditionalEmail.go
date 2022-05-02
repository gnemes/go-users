package serializerentities

import (
	"errors"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

type AdditionalEmail struct {
	*BaseSerializerEntity `json:"-"`
	
	Email     string  `json:"email"`
}

func (se *AdditionalEmail) GetName() string {
	return jsonapitypes.JsonapiAdditionalEmailType
}

func (se *AdditionalEmail) Fill(e entities.Entity) error {
	if additionalEmail, ok := e.(*entities.AdditionalEmail); ok {
		se.Email = additionalEmail.Email
		se.SetID(additionalEmail.ID)
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}