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
	http.Handle("/list", http.HandlerFunc(us.list))
	http.HandleFunc("/user", us.user)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("connect ok")
}

// HandleFunc registers the handler function for the given pattern.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	mux.Handle(pattern, HandlerFunc(handler))
// }

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.

// func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.

// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
// 	DefaultServeMux.HandleFunc(pattern, handler)
// }
