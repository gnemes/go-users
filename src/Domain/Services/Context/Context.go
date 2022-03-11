package context

import (
	logger "github.com/gnemes/go-users/Domain/Services/Logger"
)

type Context struct {
	Logger logger.Logger
	Items  map[string]interface{}
}

func (c *Context) Get(key string) interface{} {
	if value, ok := c.Items[key]; ok {
		return value
	} else {
		return nil
	}
}

func NewQuery() *Context {
	var Items map[string]interface{}

	Items = make(map[string]interface{})

	newContext := Context{
		Items: Items,
	}

	return &newContext
}