package entities

import (

)

type Platform struct {
	ID   string
	Name string
}

func (e *Platform) GetID() string {
	return e.ID
}

func (e *Platform) Validate() error {
	return nil
}