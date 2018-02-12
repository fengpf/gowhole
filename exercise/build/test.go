package main

import "C"

//export Hello
func Hello(s string) {
	println(s)
}

func main() {
	Hello("sss")
}

//go build -x -v -ldflags "-s -w" -buildmode=c-shared  -o libtest.so test.go
