package usecases

import (
	//jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
)

type OutputPort interface {
	SetData(data interface{})
	// SetMetadata(m jsonapi.Meta)
}