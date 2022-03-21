package didependencies

import (
	di "github.com/sarulabs/di/v2"
	
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	presenters "github.com/gnemes/go-users/Infrastructure/Controller/Presenters"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
	usecases "github.com/gnemes/go-users/Application/UseCases"
)

var Presenters = []di.Def{
	{
		Name:  "AdminGetUserPresenter",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &presenters.JsonApiPresenter{
				Logger:     ctn.Get("Logger").(logger.Logger),
				OutputPort: ctn.Get("AdminGetUserOutputPort").(*usecases.SingleOutputPort),
				Serializer: ctn.Get("UserSerializer").(*serializers.Serializer),
			}, nil
		},
	},
}