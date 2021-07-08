package server

import (
	"net/http"
)

// Handler http request handler.
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// HandlerFunc http request handler function.
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func httpFunc(method, pattern string, handlers ...HandlerFunc) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		handler(method, w, r, handlers...)
	})
}

func handler(method string, w http.ResponseWriter, r *http.Request, handlers ...HandlerFunc) {
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	for _, h := range handlers {
		// spew.Dump(r.URL.RawQuery)
		c.Request = r
		c.ResponseWriter = w
		h.ServeHTTP(w, r)
		if err := c.Err(); err != nil {
			return
		}
	}
}
