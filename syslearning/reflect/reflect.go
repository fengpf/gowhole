package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)
	fmt.Println(v, t)
}
