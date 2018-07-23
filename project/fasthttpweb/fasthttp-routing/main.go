package main

import (
	"fmt"

	_ "net/http/pprof"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {
	router := routing.New()
	router.Get("/", func(c *routing.Context) error {
		fmt.Fprintf(c, "Hello "+c.RemoteAddr().String())
		return nil
	})
	panic(fasthttp.ListenAndServe(":8000", router.HandleRequest))
}
