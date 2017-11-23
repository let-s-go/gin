package gin

import (
	"strconv"
)

func (c *Context) ParamInt(key string) int {
	value, _ := strconv.Atoi(c.Param(key))
	return value
}

func (c *Context) QueryInt(key string) int {
	value, _ := strconv.Atoi(c.Query(key))
	return value
}

func (c *Context) QueryBool(key string) bool {
	if c.Query(key) == "true" {
		return true
	}
	return false
}

func (c *Context) RHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) RHeaderInt(key string) int {
	value, _ := strconv.Atoi(c.Request.Header.Get(key))
	return value
}
