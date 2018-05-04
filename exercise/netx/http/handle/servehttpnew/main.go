package main

import (
	"fmt"
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

func (u users) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for id, name := range u {
			fmt.Fprintf(w, "ID(%d),Name(%s)\n", id, name)
		}
	case "/user":
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
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such user page(%s)\n", r.URL)
	}
}

func main() {
	us := users{
		1: "tom",
		2: "jack",
	}
	http.ListenAndServe(addr, us)
}
