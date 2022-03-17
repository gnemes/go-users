package di

import (
	di "github.com/sarulabs/di/v2"

	didependencies "github.com/gnemes/go-users/Infrastructure/Services/Di/Dependencies"
)

func BuildDi() di.Container {
	builder, _ := di.NewBuilder()

	builder.Add(didependencies.Base...)
	builder.Add(didependencies.ControllersHttp...)
	builder.Add(didependencies.Middlewares...)
	builder.Add(didependencies.Presenters...)
	builder.Add(didependencies.Repositories...)
	builder.Add(didependencies.Serializers...)
	builder.Add(didependencies.UseCases...)
	
	return builder.Build()
}
