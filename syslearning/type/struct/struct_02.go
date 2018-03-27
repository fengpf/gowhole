package main

import (
	"fmt"
	"gowhole/syslearning/type/struct/aaa"
)

type Skills []string

type Student_ struct {
	aaa.Human
	//Skills
	//int
	speciality string
}

func main() {
	mark := aaa.Student{aaa.Human{"mark", 25, 170}, "Computer"}

	fmt.Println(mark.Name)
	fmt.Println(mark.Human.Name)

	//mark := Student{Human{"mark", 25, 170}, "Computer"}

	jane := Student_{aaa.Human{"mark", 25, 170}, "Computer"}

	fmt.Println(jane)
}
