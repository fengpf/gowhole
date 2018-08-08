package main

//逃逸通常是内存需要合法访问，提升生存周期，或者对象过大，不适合放在栈上，但是G.stack的上限是1GB，所以得具体问题具体分析
//同理就算是引用类型，比如slice，只要内存合法，也未必会就在堆上分配

//结论：
//逃逸还是要关注的，因为在堆上分配的对象过多，意味着GC mark的工作量会变大；
//尤其是工程量大的程序，可能会因对象过多，造成GC频繁启动；
//逃逸要在堆上分配内存，malloc gc需要调用很多函数，必然导致性能问题；
//在局部范围内使用unsafe不存在问题
func main() {
	x := []int{1, 2, 3}
	test(x)

}

func test(x []int) {
	x[1] = 100
}

//main 向test传指针，所以 test stack frame 通过指针实现访问 main stack frame 的数据

// go build -o main -gcflags "-l -m"
// # gowhole/exercise/runtime/stack/escape_analysis_ref_type
// ./main.go:10:15: test x does not escape
// ./main.go:5:12: main []int literal does not escape

// go tool objdump -s "main\.main" main
// TEXT main.main(SB) /data/app/go/src/gowhole/exercise/runtime/stack/escape_analysis_ref_type/main.go
//   main.go:4             0x104aad0               65488b0c25a0080000      MOVQ GS:0x8a0, CX
//   main.go:4             0x104aad9               483b6110                CMPQ 0x10(CX), SP
//   main.go:4             0x104aadd               7650                    JBE 0x104ab2f
//   main.go:4             0x104aadf               4883ec38                SUBQ $0x38, SP
//   main.go:4             0x104aae3               48896c2430              MOVQ BP, 0x30(SP)
//   main.go:4             0x104aae8               488d6c2430              LEAQ 0x30(SP), BP
//   main.go:5             0x104aaed               488b055c700200          MOVQ main.statictmp_0(SB), AX
//   main.go:5             0x104aaf4               4889442418              MOVQ AX, 0x18(SP)
//   main.go:5             0x104aaf9               0f100558700200          MOVUPS main.statictmp_0+8(SB), X0
//   main.go:5             0x104ab00               0f11442420              MOVUPS X0, 0x20(SP)
//   main.go:6             0x104ab05               488d442418              LEAQ 0x18(SP), AX
//   main.go:6             0x104ab0a               48890424                MOVQ AX, 0(SP)
//   main.go:6             0x104ab0e               48c744240803000000      MOVQ $0x3, 0x8(SP)
//   main.go:6             0x104ab17               48c744241003000000      MOVQ $0x3, 0x10(SP)
//   main.go:6             0x104ab20               e81b000000              CALL main.test(SB)
//   main.go:8             0x104ab25               488b6c2430              MOVQ 0x30(SP), BP
//   main.go:8             0x104ab2a               4883c438                ADDQ $0x38, SP
//   main.go:8             0x104ab2e               c3                      RET
//   main.go:4             0x104ab2f               e8ec88ffff              CALL runtime.morestack_noctxt(SB)
//   main.go:4             0x104ab34               eb9a                    JMP main.main(SB)
//   :-1                   0x104ab36               cc                      INT $0x3
//   :-1                   0x104ab37               cc                      INT $0x3
//   :-1                   0x104ab38               cc                      INT $0x3
//   :-1                   0x104ab39               cc                      INT $0x3
//   :-1                   0x104ab3a               cc                      INT $0x3
//   :-1                   0x104ab3b               cc                      INT $0x3
//   :-1                   0x104ab3c               cc                      INT $0x3
//   :-1                   0x104ab3d               cc                      INT $0x3
//   :-1                   0x104ab3e               cc                      INT $0x3
//   :-1                   0x104ab3f               cc                      INT $0x3

// go build -o main -gcflags "-m"
// # gowhole/exercise/runtime/stack/escape_analysis_ref_type
// ./main.go:10:6: can inline test
// ./main.go:4:6: can inline main
// ./main.go:6:6: inlining call to test
// ./main.go:5:12: main []int literal does not escape
// ./main.go:10:15: test x does not escape
