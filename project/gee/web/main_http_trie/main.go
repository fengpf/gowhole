package main

import (
	"gowhole/project/gee/web/engine"
	"net/http"
)

func main() {
	e := engine.New()

	e.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	e.GET("/hello", func(c *engine.Context) {
		// expect /hello?name=gee
		c.Plain(http.StatusOK,
			"hello %s, you're at %s\n",
			c.Query("name"),
			c.Path,
		)
	})

	e.GET("/hello/:name", func(c *engine.Context) {
		// expect /hello?name=gee
		c.Plain(http.StatusOK,
			"hello %s, you're at %s\n",
			c.Query("name"),
			c.Path,
		)
	})

	e.GET("/assets/*filepath", func(c *engine.Context) {
		// expect /hello?name=gee
		c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
	})

	e.Run(":9999")
}
