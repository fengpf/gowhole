package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 100
	return func() int {
		// fmt.Printf("%v\n", i)
		i += 1
		return i
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		// println(x)
		sum += x
		println(sum)
		return sum
	}
}

func p(i int) {
	println(i)
}

func main() {
	a := []int{1, 2, 3}
	for _, v := range a {
		println(v)
		// defer p(v)
		defer func() {
			fmt.Printf("%p\n", &v)
			println(v)
		}()
	}
	//We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	// nextInt := intSeq()
	//See the effect of the closure by calling nextInt a few times.
	// fmt.Println(nextInt)
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// pos, _ := adder(), adder()
	// for i := 1; i < 3; i++ {
	// 	pos(i)
	// 	// fmt.Println(
	// 	// 	pos(i),
	// 	// 	// neg(-2*i),
	// 	// )
	// }
}
