package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			msg := make(map[string]string)
			msg["message"] = "pong"
			m, _ := json.Marshal(msg)
			fmt.Fprintf(w, string(m))
		}

	})

	log.Fatal(http.ListenAndServe(":9000", nil))
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
