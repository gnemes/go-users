package entities

import ()

type UserProfile struct {
	ID       string
	Name     string
	LastName string
	Age      *int
	Phone    *string
}

func (e *UserProfile) GetID() string {
	return e.ID
}

func (e *UserProfile) Validate() error {
	return nil
}