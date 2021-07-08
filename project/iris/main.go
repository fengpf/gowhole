package main

import (
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/pprof"
)

var (
	app = iris.New()
	// app  = iris.Default()// use Logger(), Recovery()
)

func main() {

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
