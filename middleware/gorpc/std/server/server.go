package main

import "gowhole/middleware/gorpc/std/test"

func main() {
	//test.StartServerByDefault()
	//fmt.Println("start rpc by default")
	//make(chan struct{}) <- struct{}{}


	test.StartServer()
}
