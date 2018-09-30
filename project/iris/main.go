package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/pprof"
)

func main() {
	app := iris.Default()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})

	app.Any("/debug/pprof/{action:path}", pprof.New())
	//                              ___________

	app.Get("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "pong",
		})
	})
	app.Run(iris.Addr(":9000"))
}
