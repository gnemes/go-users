package didependencies

import (
	di "github.com/sarulabs/di/v2"

	context "github.com/gnemes/go-users/Domain/Services/Context"
	controllerhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http"
	queryhttp "github.com/gnemes/go-users/Infrastructure/Controller/Http/Query"
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
	serializers "github.com/gnemes/go-users/Infrastructure/Serializers"
)

var ControllersHttp = []di.Def{
	{
		Name:  "QueryHttp",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			var filters map[string]interface{}
			var includes map[string]bool
			var sorts []queryhttp.QuerySort

			filters = make(map[string]interface{})
			includes = make(map[string]bool)
			sorts = make([]queryhttp.QuerySort, 0)

			newQuery := &queryhttp.Query{
				Offset:   0,
				Limit:    0,
				Filters:  filters,
				Includes: includes,
				Sorts:    sorts,
			}

			return newQuery, nil
		},
	},
	{
		Name:  "GetUserControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Get{
				Base: ctn.Get("BaseControllerHttp").(*controllerhttp.Base),
			}, nil
		},
	},
	{
		Name:  "BaseControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Base{
				Logger:  ctn.Get("Logger").(logger.Logger),
				Context: ctn.Get("Context").(*context.Context),
			}, nil
		},
	},
	{
		Name:  "ErrorControllerHttp",
		Scope: di.Request,
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			return &controllerhttp.Error{
				Logger: ctn.Get("Logger").(logger.Logger),
				ErrorSerializer: ctn.Get("HttpErrorSerializer").(*serializers.ErrorSerializer),
			}, nil
		},
	},
}