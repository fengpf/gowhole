package context

import (
	"net/http"
	"time"
)

//Context for wx.
type Context struct {
	APPID       string
	APPSecret   string
	AccessToken string

	Request        *http.Request
	ResponseWriter http.ResponseWriter

	Keys map[string]interface{}
}

//New for wx config
func New(req *http.Request, resp http.ResponseWriter) *Context {
	return &Context{
		ResponseWriter: resp,
		Request:        req,
	}
}

//Set set an data by key-value.
func (c *Context) Set(key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

//Get get an data by key-value.
func (c *Context) Get(key string) (value interface{}, ok bool) {
	value, ok = c.Keys[key]
	return
}

// Deadline implement golang.org cxt.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Done implement golang.org cxt.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err implement golang.org cxt.
func (c *Context) Err() error {
	return nil
}

// Value implement golang.org cxt.
func (c *Context) Value(key interface{}) interface{} {
	if key == 0 {
		return c.Request
	}
	if keyAsString, ok := key.(string); ok {
		val, _ := c.Get(keyAsString)
		return val
	}
	return nil
}
