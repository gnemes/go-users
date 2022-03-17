package didependencies

import (
	di "github.com/sarulabs/di/v2"

	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	serializersentities "github.com/gnemes/go-users/Infrastructure/Serializers/Entities"
)

var Serializers = []di.Def{
	{
		Name:  "HttpErrorSerializer",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			base := &serializersentities.BaseSerializerEntity{}
			entity := &serializersentities.Error{
				BaseSerializerEntity: base,
			}
			return &serializers.Serializer{
				SerializerEntity: entity,
			}, nil
		},
	},
	{
		Name:  "UserSerializer",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			base := &serializersentities.BaseSerializerEntity{}
			entity := &serializersentities.User{
				BaseSerializerEntity: base,
			}
			return &serializers.Serializer{
				SerializerEntity: entity,
			}, nil
		},
	},
}