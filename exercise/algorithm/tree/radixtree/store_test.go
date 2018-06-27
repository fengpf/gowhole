package main

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	handlers := []string{}

	handlers = append(handlers, "a")
	add("GET", "/", handlers)

	handlers = append(handlers, "b")
	add("GET", "/test", handlers)

}

func Test_get(t *testing.T) {
	handlers := []string{}
	handlers = append(handlers, "a")
	add("GET", "/", handlers)

	h, pnames := find("GET", "/", []string{})
	fmt.Println(h, pnames)
}
