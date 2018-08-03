package main

import (
	"fmt"
	"os"

	"gowhole/exercise/basictests/dlv/lib"
)

func main() {
	fmt.Println("Golang dbg test...")
	var argc = len(os.Args)
	var argv = append([]string{}, os.Args...)

	fmt.Printf("argc:%d\n", argc)
	fmt.Printf("argv:%v\n", argv)

	var var1 = 1
	var var2 = "golang dbg test"
	var var3 = []int{1, 2, 3}
	var var4 lib.MyStruct

	var4.A = 1
	var4.B = "golang dbg my struct field B"
	var4.C = map[int]string{1: "value1", 2: "value2", 3: "value3"}
	var4.D = []string{"D1", "D2", "D3"}

	lib.DBGTestRun(var1, var2, var3, var4)
	fmt.Println("Golang dbg test over")
}
