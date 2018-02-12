package main

import (
	"fmt"
	"goplay/x/src/flag"
)

func main() {
	var myFlagSet = flag.NewFlagSet("myflagset", flag.ExitOnError)
	abc := myFlagSet.String("abc", "default value", "help mesage")
	ghi := myFlagSet.Bool("def", true, "help mesage")
	// myFlagSet.Bool("fdf", true, "help mesage")

	myFlagSet.Parse([]string{"-abc", "abc-value", "-def", "sss"})
	fmt.Println("输出的参数abc的值是:", *abc)
	fmt.Println("输出的参数ghi的值是:", *ghi)
	args := myFlagSet.Args()
	for i := range args {
		fmt.Println(i, myFlagSet.Arg(i))
	}
}
