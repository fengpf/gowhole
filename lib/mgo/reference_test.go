package main

import (
	"fmt"
	"testing"
)

type Cat struct {
	Name string
}

func A() (c Cat) {
	c.Name = "mimi"
	fmt.Println(&c)
	return
}

func B() (s string) {
	s = "hah"
	return
}

func Test_cat(t *testing.T) {
	fmt.Println(A())
	fmt.Println(B())
}
