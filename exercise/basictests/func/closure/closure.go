package main

func main() {
	println(deferFunc1(1))
	println(deferFunc2(1))
	println(deferFunc3(1))
}

func deferFunc1(i int) (t int) {
	t = i
	// println(&t)
	defer func() {
		// println(&t)
		t += 3
	}()
	return t
}

func deferFunc2(i int) int {
	t := i
	println(&t, t)
	defer func() {
		println("deferFunc2 start defer", &t, t)
		t += 3
		println("deferFunc2 stop defer", &t, t)
	}()
	println(&t, t)
	return t
}

func deferFunc3(i int) (t int) {
	println(&t, t)
	defer func() {
		println("deferFunc3 start defer", &t, t)
		t += i
		println("deferFunc3 stop defer", &t, t)
	}()
	println(&t, t)
	return 2
}
