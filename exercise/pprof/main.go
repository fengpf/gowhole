package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"

	_ "net/http/pprof"
)

//应用程序用法
//go run main.go --cpuprofile=cpu.prof --memprofile=mem.prof
//go tool pprof cpu.prof / mem.prof

//web程序用法
//引入包 _ "net/http/pprof" 就好
//运行程序，模拟web请求
//浏览器中查看 http://localhost:9000/debug/pprof/
//go tool pprof http://localhost:9000/debug/pprof/profile
//输入命令 web，找到生成svg文件浏览器打开
var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile = flag.String("memprofile", "", "write memory profile to `file`")
)

func init() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// ... rest of the program ...

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}

}

func main() {
	// 应用程序测试代码
	// var a [10]int
	// for i := 0; i < 4; i++ {
	// 	go func(i int) {
	// 		for {
	// 			a[i]++
	// 		}
	// 	}(i)
	// }
	// time.Sleep(time.Millisecond)
	// fmt.Println(a)

	//web程序测试代码
	// m := http.NewServeMux()
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
