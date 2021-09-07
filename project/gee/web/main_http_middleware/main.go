package main

import (
	"log"
	"net/http"
	"time"

	"gowhole/project/gee/web/engine"
)

func onlyForV2() engine.HandlerFunc {
	return func(c *engine.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := engine.New()
	r.Use(engine.Logger()) // global midlleware
	r.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello engine</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *engine.Context) {
			// expect /hello/enginektutu
			c.Plain(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
