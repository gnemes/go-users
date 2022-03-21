package usecases

import (
	entities "github.com/gnemes/go-users/Domain/Model/Entities"
	jsonapi "github.com/gnemes/go-users/Infrastructure/Services/Jsonapi"
)

type MultipleOutputPort struct {
	Entities []entities.Entity
	Metadata jsonapi.Meta
	Totals   *uint
}

func (op *MultipleOutputPort) SetData(data interface{}) {
	op.Entities = data.([]entities.Entity)
}

func (op *MultipleOutputPort) SetMetadata(m jsonapi.Meta) {
	op.Metadata = m
}

func (op *MultipleOutputPort) SetTotals(t *uint) {
	op.Totals = t
}

func (op *MultipleOutputPort) GetData() interface{} {
	return op.Entities
}