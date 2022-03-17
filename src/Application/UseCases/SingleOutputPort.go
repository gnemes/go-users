package usecases

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
)

type SingleOutputPort struct {
	Entity   entities.Entity
	Metadata jsonapi.Meta
}

func (op *SingleOutputPort) SetData(data interface{}) {
	op.Entity = data.(entities.Entity)
}

func (op *SingleOutputPort) SetMetadata(m jsonapi.Meta) {
	op.Metadata = m
}

func (op *SingleOutputPort) GetData() interface{} {
	return op.Entity
}