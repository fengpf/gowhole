package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

const (
	addr = "127.0.0.1:8000"
)

type users map[int]string

func (u users) list(w http.ResponseWriter, r *http.Request) {
	for id, name := range u {
		fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
	}
}

func (u users) user(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Fprintf(w, "strconv.ParseInt id(%s)|error(%v)\n", idStr, err)
	}
	name, ok := u[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such user id(%d)|error(%v)\n", id, err)
		return
	}
	fmt.Fprintf(w, "user id(%d)|name(%s)\n", id, name)
}

func main() {
	us := users{
		1: "tom",
		2: "jack",
	}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(us.list)) //语句http.HandlerFunc(db.list)是一个转换而非一个函数调用，因为http.HandlerFunc是一个类型
	mux.Handle("/user", http.HandlerFunc(us.user))
	log.Fatal(http.ListenAndServe(addr, mux))
}

// type HandlerFunc func(ResponseWriter, *Request)

//ServeHTTP calls f(w, r).
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
// 	f(w, r)
// }
