package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

func main() {

	go dummyAllocations()
	go dummyCPUUsage()

	go func() {
		router := routing.New()
		router.Get("/", func(c *routing.Context) error {
			fmt.Fprintf(c, "Hello, world!")
			return nil
		})
		panic(fasthttp.ListenAndServe(":8000", router.HandleRequest))
	}()
	log.Printf("you can now open http://localhost:8000/debug/charts/ in your browser")
	select {}
}

func dummyCPUUsage() {
	var a uint64
	var t = time.Now()
	for {
		t = time.Now()
		a += uint64(t.Unix())
	}
}

func dummyAllocations() {
	var d []uint64

	for {
		for i := 0; i < 2*1024*1024; i++ {
			d = append(d, 42)
		}
		time.Sleep(time.Second * 10)
		fmt.Println(len(d))
		d = make([]uint64, 0)
		runtime.GC()
		time.Sleep(time.Second * 10)
	}
}
