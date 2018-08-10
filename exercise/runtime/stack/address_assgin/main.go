package main

// go build -o main -gcflags -N
// go tool objdump -S -s "main\.main" main

// go tool compile -N -l -S main.go > main.s

func main() {
	var (
		a struct{}
		b int
		c = new(struct{})
		d = new([0]int)
		e bool
		f byte
	)
	_ = a
	_ = b
	_ = c
	_ = d
	_ = e
	_ = f
	// println(
	// 	&a, //0xc42004bf38  a和d地址相同 分配在堆上内存地址共享了
	// 	&b, //0xc42004bf40
	// 	c,  //0xc42004bf3e  c和f地址相同
	// 	d,  //0xc42004bf38
	// 	&e, //0xc42004bf3f
	// 	&f, //0xc42004bf3f
	// )
	// println(
	// 	unsafe.Sizeof(a), //0
	// 	unsafe.Sizeof(b), //8
	// 	unsafe.Sizeof(c), //8
	// 	unsafe.Sizeof(d), //8
	// 	unsafe.Sizeof(e), //1
	// 	unsafe.Sizeof(f), //1
	// )
}
