package main

import (
	"gowhole/project/gee/web/engine"
	"net/http"
)

func main()  {
	e:=engine.New()

	e.GET("/index", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	//curl "http://localhost:9999/v1/hello?name=geektutu"
	v1 := e.Group("/v1")
	{
		v1.GET("/", func(c *engine.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *engine.Context) {
			// expect /hello?name=geektutu
			c.Plain(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := e.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *engine.Context) {
			// expect /hello/geektutu
			c.Plain(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *engine.Context) {
			c.JSON(http.StatusOK, engine.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	e.Run(":9999")
}
