package context

import (
	"errors"

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

func (c *Context) Add(key string, value interface{}) error {
	if _, ok := c.Items[key]; ok {
		return errors.New("Item already exists in context.")
	} 

	c.Items[key] = value
	return nil
}