package serializerentities

import (
	"errors"
	"strconv"
	"encoding/json"

	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapitypes "github.com/gnemes/go-users/Infrastructure/Serializers/JsonapiTypes"
)

type ErrorContainer struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	*BaseSerializerEntity                       `json:"-"`
	Status                string                `json:"status"`
	Code                  string                `json:"code"`
	Title                 string                `json:"title"`
	Detail                string                `json:"detail"`
}

func (se *Error) GetName() string {
	return jsonapitypes.JsonapiErrorType
}

func (se *Error) Fill(e entities.Entity) error {
	if errEntity, ok := e.(*entities.Error); ok {
		se.Status = strconv.Itoa(errEntity.Status)
		se.Code   = errEntity.Code
		se.Title  = errEntity.Title
		se.Detail = errEntity.Detail
		se.SetID(errEntity.ID)
	} else {
		return errors.New("Invalid entity for this serializer")
	}

	return nil
}

func (se *Error) SerializeEntity(e entities.Entity) ([]byte, error) {
	var errors []*Error
	
	err := se.Fill(e)
	if err != nil {
		return nil, err
	}

	errors = append(errors, se)
	
	errorContainer := ErrorContainer{
		Errors: errors,
	}
	
	return json.Marshal(errorContainer)
}