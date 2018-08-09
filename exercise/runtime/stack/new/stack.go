package main

func main() {
	a, b := 1, 2
	_ = add1(a, b)
	_ = add2(a, b)
}
func add1(x, y int) int {
	return x + y
}
func add2(x, y int) int {
	_ = make([]byte, 200)
	return x + y
}

// FP: Frame pointer: arguments and locals.
// PC: Program counter: jumps and branches.
// SB: Static base pointer: global symbols.
// SP: Stack pointer: top of stack.

// All user-defined symbols are written as offsets to the pseudo-registers FP (arguments and locals) and SB (globals).
