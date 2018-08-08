package main

func main() {
	p := test() //如果是内联的话，name都是main stack frame，自然内存合法，不需要逃逸
	println(*p)
}

func test() *int {
	x := 0x100
	return &x //ret 以后，其stack frame就是不合法的，所以必须放到堆上
}

// go build -o main -gcflags "-l -m"

// # gowhole/exercise/runtime/stack/escape_analysis
// ./main.go:10:9: &x escapes to heap
// ./main.go:9:2: moved to heap: x

// go build -o main -gcflags "-m"
// # gowhole/exercise/runtime/stack/escape_analysis
// ./main.go:8:6: can inline test
// ./main.go:3:6: can inline main
// ./main.go:4:11: inlining call to test
// ./main.go:4:11: main &x does not escape
// ./main.go:10:9: &x escapes to heap
// ./main.go:9:2: moved to heap: x

// go tool objdump -S -s "main\.test" main
// TEXT main.test(SB) /data/app/go/src/gowhole/exercise/runtime/stack/escape_analysis/main.go
// func test() *int {
//   0x104ab30             65488b0c25a0080000      MOVQ GS:0x8a0, CX
//   0x104ab39             483b6110                CMPQ 0x10(CX), SP
//   0x104ab3d             7639                    JBE 0x104ab78
//   0x104ab3f             4883ec18                SUBQ $0x18, SP
//   0x104ab43             48896c2410              MOVQ BP, 0x10(SP)
//   0x104ab48             488d6c2410              LEAQ 0x10(SP), BP
//         x := 0x100
//   0x104ab4d             488d05cc980000          LEAQ type.*+38976(SB), AX
//   0x104ab54             48890424                MOVQ AX, 0(SP)
//   0x104ab58             e8c303fcff              CALL runtime.newobject(SB)
//   0x104ab5d             488b442408              MOVQ 0x8(SP), AX
//   0x104ab62             48c70000010000          MOVQ $0x100, 0(AX)
//         return &x
//   0x104ab69             4889442420              MOVQ AX, 0x20(SP)
//   0x104ab6e             488b6c2410              MOVQ 0x10(SP), BP
//   0x104ab73             4883c418                ADDQ $0x18, SP
//   0x104ab77             c3                      RET
// func test() *int {
//   0x104ab78             e8a388ffff              CALL runtime.morestack_noctxt(SB)
//   0x104ab7d             ebb1                    JMP main.test(SB)

//   0x104ab7f             cc                      INT $0x3
