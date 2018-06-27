package main

import "fmt"

var (
	stores map[string]routeStore
)

// routeStore stores route paths and the corresponding handlers.
type routeStore interface {
	Add(key string, data interface{}) int
	Get(key string, pvalues []string) (data interface{}, pnames []string)
	String() string
}

func add(method, path string, handlers []string) {
	stores = make(map[string]routeStore)
	store := stores[method]
	if store == nil {
		store = newStore()
		stores[method] = store
	}
	store.Add(path, handlers)
}

func find(method, path string, pvalues []string) (handlers []string, pnames []string) {
	var hh interface{}
	if store := stores[method]; store != nil {
		hh, pnames = store.Get(path, pvalues)
	}
	if hh != nil {
		return hh.([]string), pnames
	}
	return []string{}, pnames
}

func main() {
	handlers := []string{}

	handlers = append(handlers, "a")
	add("GET", "/test/add", handlers)

	h, pnames := find("GET", "/", []string{})
	fmt.Println(h, pnames)
}
