package serializers

import (
	"encoding/json"
	"strconv"
	
	httperror "github.com/gnemes/go-users/Infrastructure/Controller/Http/Error"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
)

// ErrorContainerSerializer
type ErrorContainerSerializer struct {
	Errors []ErrorSerializer `json:"errors"`
}

// ErrorSerializer struct
type ErrorSerializer struct {
	ID          string    `json:"-"`
	Status      string    `json:"status"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Detail      string    `json:"detail"`

	Uuid        uuid.Uuid `json:"-"`
}

/*
Error{
	ID: "001",
	Links: &ErrorLinks{
		About: "http://bla/blub",
	},
	Status: "500",
	Code:   "001",
	Title:  "Title must not be empty",
	Detail: "Never occures in real life",
	Source: &ErrorSource{
		Pointer: "#titleField",
	},
	Meta: map[string]interface{}{
		"creator": "api2go",
	},
}
*/

// GetName serializer name
func (e ErrorSerializer) GetName() string {
	return "errors"
}

// GetID return id of resource
func (e ErrorSerializer) GetID() string {
	return e.ID
}

// SetID
func (e *ErrorSerializer) SetID(id string) error {
	e.ID = id
	return nil
}

// Serialize error in json api standar
func (e *ErrorSerializer) Serialize(er httperror.HttpError) ([]byte, error) {
	var errors []ErrorSerializer
	
	err := ErrorSerializer{
		ID: e.Uuid.New(),
		Status: strconv.Itoa(er.Status),
		Code: er.Code,
		Title: er.Title,
		Detail: er.Detail,
	}

	errors = append(errors, err)

	errorContainer := ErrorContainerSerializer{
		Errors: errors,
	}
	
	return json.Marshal(errorContainer)
}