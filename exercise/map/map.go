package main

func main() {
	a := make(map[int]int)
	a[1] = 1
	a[2] = 2
	b := a[3]
	println(b)
}
