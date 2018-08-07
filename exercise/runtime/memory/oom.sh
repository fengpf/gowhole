
runtime stack:
runtime.throw(0x4b6b4b, 0x16)
	/usr/local/go/src/runtime/panic.go:616 +0x81
runtime.sysMap(0xc420100000, 0x42b50000, 0x7f2900000000, 0x547d18)
	/usr/local/go/src/runtime/mem_linux.go:216 +0x20a
runtime.(*mheap).sysAlloc(0x52f7a0, 0x42b50000, 0x7f29935f41d0)
	/usr/local/go/src/runtime/malloc.go:470 +0xd4
runtime.(*mheap).grow(0x52f7a0, 0x215a1, 0x0)
	/usr/local/go/src/runtime/mheap.go:907 +0x60
runtime.(*mheap).allocSpanLocked(0x52f7a0, 0x215a1, 0x547d28, 0x7ffec488f390)
	/usr/local/go/src/runtime/mheap.go:820 +0x301
runtime.(*mheap).alloc_m(0x52f7a0, 0x215a1, 0x410101, 0xc41fffc8ff)
	/usr/local/go/src/runtime/mheap.go:686 +0x118
runtime.(*mheap).alloc.func1()
	/usr/local/go/src/runtime/mheap.go:753 +0x4d
runtime.(*mheap).alloc(0x52f7a0, 0x215a1, 0xc420000101, 0x52b460)
	/usr/local/go/src/runtime/mheap.go:752 +0x8a
runtime.largeAlloc(0x42b42000, 0x540100, 0x7f29935931c0)
	/usr/local/go/src/runtime/malloc.go:826 +0x94
runtime.mallocgc.func1()
	/usr/local/go/src/runtime/malloc.go:721 +0x46
runtime.systemstack(0x0)
	/usr/local/go/src/runtime/asm_amd64.s:409 +0x79
runtime.mstart()
	/usr/local/go/src/runtime/proc.go:1175

goroutine 1 [running]:
runtime.systemstack_switch()
	/usr/local/go/src/runtime/asm_amd64.s:363 fp=0xc420041598 sp=0xc420041590 pc=0x44c190
runtime.mallocgc(0x42b42000, 0x0, 0xc420000100, 0x300000002)
	/usr/local/go/src/runtime/malloc.go:720 +0x8a2 fp=0xc420041638 sp=0xc420041598 pc=0x40dbb2
runtime.growslice(0x494b00, 0xc42006c000, 0x37, 0x40, 0x42b40037, 0xc4200416f8, 0x40bf03, 0x49d520)
	/usr/local/go/src/runtime/slice.go:172 +0x21d fp=0xc4200416a0 sp=0xc420041638 pc=0x43971d
fmt.(*buffer).WriteString(...)
	/usr/local/go/src/fmt/print.go:82
fmt.(*fmt).padString(0xc42006a040, 0x12, 0x42b40000)
	/usr/local/go/src/fmt/format.go:110 +0x107 fp=0xc420041728 sp=0xc4200416a0 pc=0x479887
fmt.(*fmt).fmt_s(0xc42006a040, 0x12, 0x42b40000)
	/usr/local/go/src/fmt/format.go:328 +0x61 fp=0xc420041760 sp=0xc420041728 pc=0x47a5c1
fmt.(*pp).fmtString(0xc42006a000, 0x12, 0x42b40000, 0x76)
	/usr/local/go/src/fmt/print.go:437 +0x11f fp=0xc420041798 sp=0xc420041760 pc=0x47d63f
fmt.(*pp).printArg(0xc42006a000, 0x4949c0, 0xc42000e220, 0x76)
	/usr/local/go/src/fmt/print.go:671 +0x789 fp=0xc420041810 sp=0xc420041798 pc=0x47f599
fmt.(*pp).doPrintln(0xc42006a000, 0xc420041be0, 0x6, 0x6)
	/usr/local/go/src/fmt/print.go:1146 +0x45 fp=0xc420041880 sp=0xc420041810 pc=0x483825
fmt.Fprintln(0x4c5da0, 0xc42000c020, 0xc420041be0, 0x6, 0x6, 0x4949c0, 0x426a01, 0xc42000e220)
	/usr/local/go/src/fmt/print.go:254 +0x58 fp=0xc4200418e8 sp=0xc420041880 pc=0x47c048
fmt.Println(0xc420041be0, 0x6, 0x6, 0xc42000e220, 0x0, 0x0)
	/usr/local/go/src/fmt/print.go:264 +0x57 fp=0xc420041938 sp=0xc4200418e8 pc=0x47c147
main.main()
	/data/go/src/pub/unsafe/main.go:73 +0xdfb fp=0xc420041f88 sp=0xc420041938 pc=0x484a5b
runtime.main()
	/usr/local/go/src/runtime/proc.go:198 +0x212 fp=0xc420041fe0 sp=0xc420041f88 pc=0x427992
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:2361 +0x1 fp=0xc420041fe8 sp=0xc420041fe0 pc=0x44e621