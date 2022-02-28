package di

import (
	"sync"

	"github.com/sarulabs/di"
)

var once sync.Once
var instance *Container

// Container for di in app
type Container struct {
	Builder   *di.Builder
	Container di.Container
}

// GetInstance and initialize container singleton instance
func GetInstance() *Container {

	once.Do(func() {
		BuildContainer()
	})

	return instance
}

// DeleteInstance delete container instance
func DeleteInstance() {
	instance.Container.Delete()
}

// Get returns an interface that can be cast afterwards.
// If the object can not be created, the Get function panics.
// obj := ctn.Get("my-object").(*MyObject)
func (c *Container) Get(key string) interface{} {
	return c.Container.Get(key)
}

// BuildContainer instance
func BuildContainer() {
	builder, _ := di.NewBuilder()

	builder.Add(Base...)
	// builder.Add(Clients...)
	builder.Add(Controllers...)
	// builder.Add(Presenters...)
	// builder.Add(Repositories...)
	// builder.Add(Serializers...)
	// builder.Add(Services...)
	// builder.Add(UseCases...)

	container := new(Container)
	container.Builder = builder
	container.Container = builder.Build()

	instance = container
}
