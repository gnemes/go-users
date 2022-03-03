package di

import (
	"github.com/sarulabs/di"

	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	uuid "github.com/gnemes/go-users/Domain/Services/Uuid"
)

var Serializers = []di.Def{
	{
		Name:  "HttpErrorSerializer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return &serializers.ErrorSerializer{
				Uuid: ctn.Get("Uuid").(uuid.Uuid),
			}, nil
		},
	},
}