package serializerentities

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
)

type BaseSerializerEntity struct {
	ID string `json:"-"`
}

func (bse *BaseSerializerEntity) GetID() string {
	return bse.ID
}

func (bse *BaseSerializerEntity) SetID(id string) error {
	bse.ID = id
	return nil
}

func (bse *BaseSerializerEntity) SerializeEntity(e entities.Entity) ([]byte, error) {
	return nil, nil
}