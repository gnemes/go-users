package entities

import (

)

type Platform struct {
	ID   string
	Name string
}

func (e *Platform) Validate() error {
	return nil
}