package entities

import (

)

type User struct {
	ID         string
	Username   string
	PlatformID string
}

func (e *User) Validate() error {
	return nil
}