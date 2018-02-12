package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gosrc/x/src/net/http"
	_ "gosrc/x/src/net/http/pprof"
)

// ResponseWriter： 生成Response的接口

// Handler： 处理请求和生成返回的接口

// ServeMux： 路由，后面会说到ServeMux也是一种Handler

// Conn : 网络连接

func main() {
	// client()
	http.HandleFunc("/test/add", testHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := r.Form
	fmt.Fprintf(w, "hello world!")
	fmt.Fprintf(w, params.Get("aid"))
}
func client() {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%v\n", string(body))
}
