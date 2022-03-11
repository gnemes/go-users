package uuid

import (
	googleuuid "github.com/google/uuid"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
)

type GoogleUuid struct {
}

func (service *GoogleUuid) New() string {
	u2 := googleuuid.New()
	return u2.String()
}

func NewUuid() (uuid.Uuid, error) {
	return &GoogleUuid{}, nil
}