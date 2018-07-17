package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mkevac/debugcharts"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

const (
	addr = "127.0.0.1:8000"
)

type users map[int]string

func (u users) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for id, name := range u {
		fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
	}
}

func main() {
	// us := users{
	// 	1: "tom",
	// 	2: "jack",
	// }
	go func() {
		log.Println(http.ListenAndServe(addr, nil))
	}()
	select {}
}
