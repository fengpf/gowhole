package main

func main() {
	x := 0x100
	test(&x)

}

func test(x *int) {
	println(*x)
}

//我们知道一个函数只能直接访问它的（函数栈）空间，或者通过（函数栈空间内的）指针，
//通过跳转访问（函数栈空间外的）外部内存。
//这意味着访问逃逸到堆上的值也需要通过指针跳转

//main 向test传指针，所以 test stack frame 通过指针实现访问 main stack frame 的数据

// go build -o main -gcflags "-l -m"
// # gowhole/exercise/runtime/stack/escape_analysis_inline
// ./main.go:9:14: test x does not escape
// ./main.go:5:7: main &x does not escape

// go tool objdump -s "main\.main" main
// TEXT main.main(SB) /data/app/go/src/gowhole/exercise/runtime/stack/escape_analysis_inline/main.go
//   main.go:3             0x104aad0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
//   main.go:3             0x104aad9               483b6110                CMPQ 0x10(CX), SP
//   main.go:3             0x104aadd               762f                    JBE 0x104ab0e
//   main.go:3             0x104aadf               4883ec18                SUBQ $0x18, SP
//   main.go:3             0x104aae3               48896c2410              MOVQ BP, 0x10(SP)
//   main.go:3             0x104aae8               488d6c2410              LEAQ 0x10(SP), BP
//   main.go:4             0x104aaed               48c744240800010000      MOVQ $0x100, 0x8(SP)
//   main.go:5             0x104aaf6               488d442408              LEAQ 0x8(SP), AX
//   main.go:5             0x104aafb               48890424                MOVQ AX, 0(SP)
//   main.go:5             0x104aaff               e81c000000              CALL main.test(SB)
//   main.go:7             0x104ab04               488b6c2410              MOVQ 0x10(SP), BP
//   main.go:7             0x104ab09               4883c418                ADDQ $0x18, SP
//   main.go:7             0x104ab0d               c3                      RET
//   main.go:3             0x104ab0e               e80d89ffff              CALL runtime.morestack_noctxt(SB)
//   main.go:3             0x104ab13               ebbb                    JMP main.main(SB)
//   :-1                   0x104ab15               cc                      INT $0x3
//   :-1                   0x104ab16               cc                      INT $0x3
//   :-1                   0x104ab17               cc                      INT $0x3
//   :-1                   0x104ab18               cc                      INT $0x3
//   :-1                   0x104ab19               cc                      INT $0x3
//   :-1                   0x104ab1a               cc                      INT $0x3
//   :-1                   0x104ab1b               cc                      INT $0x3
//   :-1                   0x104ab1c               cc                      INT $0x3
//   :-1                   0x104ab1d               cc                      INT $0x3
//   :-1                   0x104ab1e               cc                      INT $0x3
//   :-1                   0x104ab1f               cc                      INT $0x3

// go build -o main -gcflags "-m"
// # gowhole/exercise/runtime/stack/escape_analysis_inline
// ./main.go:9:6: can inline test
// ./main.go:3:6: can inline main
// ./main.go:5:6: inlining call to test
// ./main.go:5:7: main &x does not escape
// ./main.go:9:14: test x does not escape
