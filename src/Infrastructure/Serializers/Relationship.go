package serializers

import (
	serializersdomain "github.com/gnemes/go-users/Domain/Serializers"
)

type Relationship struct {
	Type        string
	Name        string
	IsNotLoaded bool
	Entity      serializersdomain.SerializerEntity
}