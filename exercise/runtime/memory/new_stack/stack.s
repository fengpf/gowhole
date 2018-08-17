"".main STEXT size=112 args=0x0 locals=0x30
    // 栈大小为48，无参数
	0x0000 00000 (stack.go:3)	TEXT	"".main(SB), $48-0
	// 通过thread local storage获取当前g(g为goroutine的的数据结构)
	0x0000 00000 (stack.go:3)	MOVQ	(TLS), CX
	// 比较SP和g.stackguard0
	0x0009 00009 (stack.go:3)	CMPQ	SP, 16(CX)
	// 小于g.stackguard0，jump到105执行栈的扩容
	0x000d 00013 (stack.go:3)	JLS	105
	// 继续执行
	0x000f 00015 (stack.go:3)	SUBQ	$48, SP
	0x0013 00019 (stack.go:3)	MOVQ	BP, 40(SP)
	0x0018 00024 (stack.go:3)	LEAQ	40(SP), BP
    // 用于垃圾回收
	0x001d 00029 (stack.go:3)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (stack.go:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001d 00029 (stack.go:4)	MOVQ	$1, "".a+32(SP)
	0x0026 00038 (stack.go:4)	MOVQ	$2, "".b+24(SP)
	// 将a放入AX寄存器
	0x002f 00047 (stack.go:5)	MOVQ	"".a+32(SP), AX
	// 参数a压栈
	0x0034 00052 (stack.go:5)	MOVQ	AX, (SP)
	// 将b放入AX寄存器
	0x0038 00056 (stack.go:5)	MOVQ	"".b+24(SP), AX
	// 参数b压栈
	0x003d 00061 (stack.go:5)	MOVQ	AX, 8(SP)
	0x0042 00066 (stack.go:5)	PCDATA	$0, $0
	// 调用add1
	0x0042 00066 (stack.go:5)	CALL	"".add1(SB)
	// 将a放入AX寄存器
	0x0047 00071 (stack.go:6)	MOVQ	"".a+32(SP), AX
	// 参数a压栈
	0x004c 00076 (stack.go:6)	MOVQ	AX, (SP)
	// 将b放入AX寄存器
	0x0050 00080 (stack.go:6)	MOVQ	"".b+24(SP), AX
	// 参数b压栈
	0x0055 00085 (stack.go:6)	MOVQ	AX, 8(SP)
	0x005a 00090 (stack.go:6)	PCDATA	$0, $0
	// 调用add2
	0x005a 00090 (stack.go:6)	CALL	"".add2(SB)
	0x005f 00095 (stack.go:7)	MOVQ	40(SP), BP
	0x0064 00100 (stack.go:7)	ADDQ	$48, SP
	0x0068 00104 (stack.go:7)	RET
	0x0069 00105 (stack.go:7)	NOP
	0x0069 00105 (stack.go:3)	PCDATA	$0, $-1
	// 调用runtime.morestack_noctxt执行栈扩容
	0x0069 00105 (stack.go:3)	CALL	runtime.morestack_noctxt(SB)
	// 返回到函数开始处继续执行
	0x006e 00110 (stack.go:3)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 5a 48  eH..%....H;a.vZH
	0x0010 83 ec 30 48 89 6c 24 28 48 8d 6c 24 28 48 c7 44  ..0H.l$(H.l$(H.D
	0x0020 24 20 01 00 00 00 48 c7 44 24 18 02 00 00 00 48  $ ....H.D$.....H
	0x0030 8b 44 24 20 48 89 04 24 48 8b 44 24 18 48 89 44  .D$ H..$H.D$.H.D
	0x0040 24 08 e8 00 00 00 00 48 8b 44 24 20 48 89 04 24  $......H.D$ H..$
	0x0050 48 8b 44 24 18 48 89 44 24 08 e8 00 00 00 00 48  H.D$.H.D$......H
	0x0060 8b 6c 24 28 48 83 c4 30 c3 e8 00 00 00 00 eb 90  .l$(H..0........
	rel 5+4 t=16 TLS+0
	rel 67+4 t=8 "".add1+0
	rel 91+4 t=8 "".add2+0
	rel 106+4 t=8 runtime.morestack_noctxt+0
"".add1 STEXT nosplit size=25 args=0x18 locals=0x0
// 栈大小为0，参数为24字节, 栈帧小于StackSmall不进行栈空间判断直接执行
	0x0000 00000 (stack.go:8)	TEXT	"".add1(SB), NOSPLIT, $0-24
	0x0000 00000 (stack.go:8)	FUNCDATA	$0, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x0000 00000 (stack.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (stack.go:8)	MOVQ	$0, "".~r2+24(SP)
	0x0009 00009 (stack.go:9)	MOVQ	"".x+8(SP), AX
	0x000e 00014 (stack.go:9)	ADDQ	"".y+16(SP), AX
	0x0013 00019 (stack.go:9)	MOVQ	AX, "".~r2+24(SP)
	0x0018 00024 (stack.go:9)	RET
	0x0000 48 c7 44 24 18 00 00 00 00 48 8b 44 24 08 48 03  H.D$.....H.D$.H.
	0x0010 44 24 10 48 89 44 24 18 c3                       D$.H.D$..
"".add2 STEXT size=148 args=0x18 locals=0xd0
    // 栈大小为208字节，参数为24字节
	0x0000 00000 (stack.go:11)	TEXT	"".add2(SB), $208-24
	// 获取当前g
	0x0000 00000 (stack.go:11)	MOVQ	(TLS), CX
	// 栈大小大于StackSmall, 计算 SP - FramSzie + StackSmall 并放入AX寄存器
	0x0009 00009 (stack.go:11)	LEAQ	-80(SP), AX
	// 比较上面计算出来的值和g.stackguard0
	0x000e 00014 (stack.go:11)	CMPQ	AX, 16(CX)
	// 小于g.stackguard0, jump到138执行栈的扩容
	0x0012 00018 (stack.go:11)	JLS	138
	// 继续执行
	0x0014 00020 (stack.go:11)	SUBQ	$208, SP
	0x001b 00027 (stack.go:11)	MOVQ	BP, 200(SP)
	0x0023 00035 (stack.go:11)	LEAQ	200(SP), BP
	0x002b 00043 (stack.go:11)	FUNCDATA	$0, gclocals·54241e171da8af6ae173d69da0236748(SB)
	0x002b 00043 (stack.go:11)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x002b 00043 (stack.go:11)	MOVQ	$0, "".~r2+232(SP)
	0x0037 00055 (stack.go:12)	MOVQ	$0, ""..autotmp_3(SP)
	0x003f 00063 (stack.go:12)	LEAQ	""..autotmp_3+8(SP), DI
	0x0044 00068 (stack.go:12)	XORPS	X0, X0
	0x0047 00071 (stack.go:12)	DUFFZERO	$247
	0x005a 00090 (stack.go:12)	LEAQ	""..autotmp_3(SP), AX
	0x005e 00094 (stack.go:12)	TESTB	AL, (AX)
	0x0060 00096 (stack.go:12)	JMP	98
	0x0062 00098 (stack.go:13)	MOVQ	"".x+216(SP), AX
	0x006a 00106 (stack.go:13)	ADDQ	"".y+224(SP), AX
	0x0072 00114 (stack.go:13)	MOVQ	AX, "".~r2+232(SP)
	0x007a 00122 (stack.go:13)	MOVQ	200(SP), BP
	0x0082 00130 (stack.go:13)	ADDQ	$208, SP
	0x0089 00137 (stack.go:13)	RET
	0x008a 00138 (stack.go:13)	NOP
	0x008a 00138 (stack.go:11)	PCDATA	$0, $-1
	// 调用runtime.morestack_noctxt完成栈扩容
	0x008a 00138 (stack.go:11)	CALL	runtime.morestack_noctxt(SB)
	// jump到函数开始的地方继续执行
	0x008f 00143 (stack.go:11)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 b0 48 3b  eH..%....H.D$.H;
	0x0010 41 10 76 76 48 81 ec d0 00 00 00 48 89 ac 24 c8  A.vvH......H..$.
	0x0020 00 00 00 48 8d ac 24 c8 00 00 00 48 c7 84 24 e8  ...H..$....H..$.
	0x0030 00 00 00 00 00 00 00 48 c7 04 24 00 00 00 00 48  .......H..$....H
	0x0040 8d 7c 24 08 0f 57 c0 48 89 6c 24 f0 48 8d 6c 24  .|$..W.H.l$.H.l$
	0x0050 f0 e8 00 00 00 00 48 8b 6d 00 48 8d 04 24 84 00  ......H.m.H..$..
	0x0060 eb 00 48 8b 84 24 d8 00 00 00 48 03 84 24 e0 00  ..H..$....H..$..
	0x0070 00 00 48 89 84 24 e8 00 00 00 48 8b ac 24 c8 00  ..H..$....H..$..
	0x0080 00 00 48 81 c4 d0 00 00 00 c3 e8 00 00 00 00 e9  ..H.............
	0x0090 6c ff ff ff                                      l...
	rel 5+4 t=16 TLS+0
	rel 82+4 t=8 runtime.duffzero+247
	rel 139+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=99 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	92
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0022 00034 (<autogenerated>:1)	CMPB	AL, $1
	0x0024 00036 (<autogenerated>:1)	JHI	40
	0x0026 00038 (<autogenerated>:1)	JMP	49
	0x0028 00040 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002c 00044 (<autogenerated>:1)	ADDQ	$8, SP
	0x0030 00048 (<autogenerated>:1)	RET
	0x0031 00049 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0038 00056 (<autogenerated>:1)	CMPB	AL, $1
	0x003a 00058 (<autogenerated>:1)	JEQ	62
	0x003c 00060 (<autogenerated>:1)	JMP	69
	0x003e 00062 (<autogenerated>:1)	PCDATA	$0, $0
	0x003e 00062 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0043 00067 (<autogenerated>:1)	UNDEF
	0x0045 00069 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x004c 00076 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0053 00083 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0057 00087 (<autogenerated>:1)	ADDQ	$8, SP
	0x005b 00091 (<autogenerated>:1)	RET
	0x005c 00092 (<autogenerated>:1)	NOP
	0x005c 00092 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005c 00092 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0061 00097 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 4d 48  eH..%....H;a.vMH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 0f b6 05 00 00  ...H.,$H.,$.....
	0x0020 00 00 3c 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08  ..<.w...H.,$H...
	0x0030 c3 0f b6 05 00 00 00 00 3c 01 74 02 eb 07 e8 00  ........<.t.....
	0x0040 00 00 00 0f 0b c6 05 00 00 00 00 01 c6 05 00 00  ................
	0x0050 00 00 02 48 8b 2c 24 48 83 c4 08 c3 e8 00 00 00  ...H.,$H........
	0x0060 00 eb 9d                                         ...
	rel 5+4 t=16 TLS+0
	rel 30+4 t=15 "".initdone·+0
	rel 52+4 t=15 "".initdone·+0
	rel 63+4 t=8 runtime.throwinit+0
	rel 71+4 t=15 "".initdone·+-1
	rel 78+4 t=15 "".initdone·+-1
	rel 93+4 t=8 runtime.morestack_noctxt+0
go.info."".main SDWARFINFO size=55
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 61 00 04 00 00 00 00 02 91 68 09 62 00 04 00  .a........h.b...
	0x0030 00 00 00 02 91 60 00                             .....`.
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+112
	rel 27+4 t=29 gofile../data/app/go/src/gowhole/exercise/runtime/stack/new/stack.go+0
	rel 36+4 t=28 go.info.int+0
	rel 47+4 t=28 go.info.int+0
go.range."".main SDWARFRANGE size=0
go.info."".add1 SDWARFINFO size=70
	0x0000 02 22 22 2e 61 64 64 31 00 00 00 00 00 00 00 00  ."".add1........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0e 78 00 00 08 00 00 00 00 01 9c 0e 79 00 00 08  .x..........y...
	0x0030 00 00 00 00 02 91 08 0e 7e 72 32 00 01 08 00 00  ........~r2.....
	0x0040 00 00 02 91 10 00                                ......
	rel 9+8 t=1 "".add1+0
	rel 17+8 t=1 "".add1+25
	rel 27+4 t=29 gofile../data/app/go/src/gowhole/exercise/runtime/stack/new/stack.go+0
	rel 37+4 t=28 go.info.int+0
	rel 48+4 t=28 go.info.int+0
	rel 62+4 t=28 go.info.int+0
go.range."".add1 SDWARFRANGE size=0
go.info."".add2 SDWARFINFO size=70
	0x0000 02 22 22 2e 61 64 64 32 00 00 00 00 00 00 00 00  ."".add2........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 0e 78 00 00 0b 00 00 00 00 01 9c 0e 79 00 00 0b  .x..........y...
	0x0030 00 00 00 00 02 91 08 0e 7e 72 32 00 01 0b 00 00  ........~r2.....
	0x0040 00 00 02 91 10 00                                ......
	rel 9+8 t=1 "".add2+0
	rel 17+8 t=1 "".add2+148
	rel 27+4 t=29 gofile../data/app/go/src/gowhole/exercise/runtime/stack/new/stack.go+0
	rel 37+4 t=28 go.info.int+0
	rel 48+4 t=28 go.info.int+0
	rel 62+4 t=28 go.info.int+0
go.range."".add2 SDWARFRANGE size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+99
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
"".initdone· SNOPTRBSS size=1
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]uint8- SRODATA dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8-+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..hashfunc200 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 c8 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memhash_varlen+0
type..eqfunc200 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 c8 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 runtime.memequal_varlen+0
type..alg200 SRODATA dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc200+0
	rel 8+8 t=1 type..eqfunc200+0
type..namedata.*[200]uint8- SRODATA dupok size=14
	0x0000 00 00 0b 2a 5b 32 30 30 5d 75 69 6e 74 38        ...*[200]uint8
type.*[200]uint8 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 78 7a 36 47 00 08 08 36 00 00 00 00 00 00 00 00  xz6G...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[200]uint8-+0
	rel 48+8 t=1 type.[200]uint8+0
runtime.gcbits. SRODATA dupok size=0
type.[200]uint8 SRODATA dupok size=72
	0x0000 c8 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 a3 66 90 a8 02 01 01 91 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 c8 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg200+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[200]uint8-+0
	rel 44+4 t=6 type.*[200]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·54241e171da8af6ae173d69da0236748 SRODATA dupok size=9
	0x0000 01 00 00 00 03 00 00 00 00                       .........
