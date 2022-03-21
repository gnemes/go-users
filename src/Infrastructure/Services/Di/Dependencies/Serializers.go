package didependencies

import (
	di "github.com/sarulabs/di/v2"

	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	serializerentities "github.com/gnemes/go-users/Infrastructure/Serializers/Entities"
)

var Serializers = []di.Def{
	{
		Name:  "HttpErrorSerializer",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			base := &serializerentities.BaseSerializerEntity{
				Container: ctn,
			}
			entity := &serializerentities.Error{
				BaseSerializerEntity: base,
			}
			return &serializers.Serializer{
				Entity: entity,
			}, nil
		},
	},
	{
		Name:  "UserSerializer",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			base := &serializerentities.BaseSerializerEntity{
				Container: ctn,
			}
			entity := &serializerentities.User{
				BaseSerializerEntity: base,
			}
			return &serializers.Serializer{
				Entity: entity,
			}, nil
		},
	},
	{
		Name:  "PlatformSerializer",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			base := &serializerentities.BaseSerializerEntity{
				Container: ctn,
			}
			entity := &serializerentities.Platform{
				BaseSerializerEntity: base,
			}
			return &serializers.Serializer{
				Entity: entity,
			}, nil
		},
	},
}