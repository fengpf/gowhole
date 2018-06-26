package main

import (
	"errors"

	"github.com/valyala/fasthttp"
	"github.com/vincentLiuxiang/lu"
)

func main() {
	app := lu.New()
	app.Use("/", func(ctx *fasthttp.RequestCtx, next func(error)) {
		ctx.SetStatusCode(200)
		next(nil)
	})

	app.Get("/api", func(ctx *fasthttp.RequestCtx, next func(error)) {
		next(errors.New("something error occour\n"))
	})

	app.Use("/test", func(ctx *fasthttp.RequestCtx, next func(error)) {
		ctx.SetBody([]byte("hello world\n"))
	})

	app.Use("/", func(err error, ctx *fasthttp.RequestCtx, next func(error)) {
		ctx.SetBody([]byte(err.Error()))
	})

	app.Listen(":8000")
}
