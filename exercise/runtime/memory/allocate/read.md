
## AlgorithmOne
`go test -run none -bench AlgorithmOne -benchtime 3s -benchmem`

`BenchmarkAlgorithmOne-4          2000000              3169 ns/op             117 B/op          2 allocs/op`

> 运行完压力测试后，我们可以看到 algOne 函数分配了两次值，每次分配了 117 个字节，每次处理需要花费3169ns。
> 这真的很棒，但我们还需要知道哪行代码造成了分配。
> 为了这个目的，我们需要生成压力测试的分析数据。

`go test -run none -bench AlgorithmOne -benchtime 3s -benchmem -memprofile mem.out`

`BenchmarkAlgorithmOne-4          2000000              3096 ns/op             117 B/op          2 allocs/op`

> 当分析内存数据时，为了轻而易举地得到我们要的信息，你会想用 -alloc_space 选项替代默认的 -inuse_space 选项。
> 这将会向你展示每一次分配发生在哪里，不管你分析数据时它是不是还在内存中。

`go tool pprof -alloc_space memcpu.test mem.out`

`(pprof) list algOne`

>         15MB   335.03MB (flat, cum)   100% of Total
>         .          .     84:
>         .          .     85:// algOne is one way to solve the problem.
>         .          .     86:func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
>         .          .     87:
>         .          .     88:   // Use a bytes Buffer to provide a stream to process.
>         .   320.03MB     89:   input := bytes.NewBuffer(data)
>         .          .     90:
>         .          .     91:   // The number of bytes we are looking for.
>         .          .     92:   size := len(find)
>         .          .     93:
>         .          .     94:   // Declare the buffers we need to process the stream.
>        15MB       15MB     95:   buf := make([]byte, size)
>        .          .     96:   end := size - 1
>        .          .     97:
>        .          .     98:   // Read in an initial number of bytes we need to get started.
>        .          .     99:   if n, err := io.ReadFull(input, buf[:end]); err != nil {
>        .          .    100:           output.Write(buf[:n])

>基于这次数据分析，我们知道了input/buf 数组在堆中的分配.
>因为 input 是指针变量，分析数据表明 input 指针变量指定的 bytes.Buffer 值分配了。
>我们先关注 input 内存分配以及弄清楚为啥会被分配。

> 我们可以假定它被分配是因为调用 bytes.NewBuffer 函数时在栈上共享了 bytes.Buffer 值。然而，
> 存在于 flat 列（pprof 输出的第一列）的值告诉我们值被分配是因为 algOne 函数共享造成了它的逃逸。
>我知道 flat 列代表在函数中的分配是因为 list 命令显示 Benchmark 函数中调用了 aglOne。

`(pprof) list Benchmark`
 >        0   335.03MB (flat, cum)   100% of Total
 >        .          .     20:
 >        .          .     21:   b.ResetTimer()
 >        .          .     22:
 >        .          .     23:   for i := 0; i < b.N; i++ {
 >        .          .     24:           output.Reset()
 >        .   335.03MB     25:           algOne(in, find, repl, &output)
 >        .          .     26:   }
 >        .          .     27:}
 >        .          .     28:
 >        .          .     29:// Capture the time it takes to execute algorithm two.
 >        .          .     30:func BenchmarkAlgorithmTwo(b *testing.B) {


>因为在 cum 列（第二列）只有一个值，这告诉我 Benchmark 没有直接分配。所有的内存分配都发生在函数调用的循环里。
>你可以看到这两个 list 调用的分配次数是匹配的。

>我们还是不知道为什么 bytes.Buffer 值被分配。这时在 go build 的时候打开 -gcflags "-m -m" 就派上用场了。
>分析数据只能告诉你哪些值逃逸，但编译命令可以告诉你为啥。

>./main.go:89:26: inlining call to bytes.NewBuffer func([]byte) *bytes.Buffer { return &bytes.Buffer literal }
>可以肯定 bytes.Buffer 值没有逃逸，因为它传递给了调用栈。这是因为没有调用 bytes.NewBuffer，函数内联处理了。


`input := bytes.NewBuffer(data)`

>因为编译器选择内联 bytes.NewBuffer 函数调用，我写的代码被转成：

`input := &bytes.Buffer{buf: data}`

>这意味着 algOne 函数直接构造 bytes.Buffer 值。
>那么，现在的问题是什么造成了值从 algOne 栈帧中逃逸？答案在我们搜索结果中的另外 5 行。

>./main.go:89:26: &bytes.Buffer literal escapes to heap
>./main.go:89:26:        from ~R0 (assign-pair) at ./main.go:89:26
>./main.go:89:26:        from input (assigned) at ./main.go:89:8
>./main.go:89:26:        from input (interface-converted) at ./main.go:99:26
>./main.go:89:26:        from input (passed to call[argument escapes]) at ./main.go:99:26

>这几行告诉我们代码中的第 99 行造成了逃逸。input 变量被赋值给一个接口变量。

```

if n, err := io.ReadFull(input, buf[:end]); err != nil {
    output.Write(buf[:n]) 
    return 
} 

```

> io.ReadFull 调用造成了接口赋值。如果你看了 io.ReadFull 函数的定义，你可以看到一个接口类型是如何接收 input 值。

``` 
type Reader interface {
    Read(p []byte) (n int, err error)
}

func ReadFull(r Reader, buf []byte) (n int, err error) {
    return ReadAtLeast(r, buf, len(buf))
}

```

>传递 bytes.Buffer 地址到调用栈，在 Reader 接口变量中存储会造成一次逃逸。
>现在我们知道使用接口变量是需要开销的：分配和重定向。
>所以，如果没有很明显的使用接口的原因，你可能不想使用接口。下面是我选择在我的代码中是否使用接口的原则。

###使用接口的情况：

>用户 API 需要提供实现细节的时候。
>API 的内部需要维护多种实现。
>可以改变的 API 部分已经被识别并需要解耦。

###不使用接口的情况：

>为了使用接口而使用接口。
>推广算法。
>当用户可以定义自己的接口时。


>现在我们可以问自己，这个算法真的需要 io.ReadFull 函数吗？答案是否定的，
>因为bytes.Buffer 类型有一个方法可以供我们使用。
>使用方法而不是调用一个函数可以防止重新分配内存。

>让我们修改代码，删除 io 包，并直接使用 Read 函数而不是 input 变量。
>修改后的代码删除了 io 包的调用，为了保留相同的行号，我使用空标志符替代 io 包的引用( _ "io" )
>这会允许（没有使用的）库导入的行待在列表中。

## AlgorithmTwo

`go test -run none -bench AlgorithmTwo -benchtime 3s -benchmem -memprofile mem.out`

`BenchmarkAlgorithmTwo-4          2000000              2144 ns/op               5 B/op          1 allocs/op`

>我们可以看到大约 32% 的性能提升。代码从 3169 ns/op 降到 2144 ns/op。 注意不用机器测试数据可能不同
>解决了这个问题，我们现在可以关注 buf 切片数组。
>如果再次使用测试代码生成分析数据，我们应该能够识别到造成剩下的分配的原因。

`go tool pprof -alloc_space memcpu.test mem.out`

>       Total: 5MB
>       5MB        5MB (flat, cum)   100% of Total
>         .          .    140:
>         .          .    141:   // The number of bytes we are looking for.
>         .          .    142:   size := len(find)
>         .          .    143:
>         .          .    144:   // Declare the buffers we need to process the stream.
>       5MB        5MB    145:   buf := make([]byte, size)
>         .          .    146:   end := size - 1
>         .          .    147:
>         .          .    148:   // Read in an initial number of bytes we need to get started.
>         .          .    149:   if n, err := input.Read(buf[:end]); err != nil || n < end {
>         .          .    150:           output.Write(buf[:n])

`go build -gcflags "-m -m"`

>./main.go:95:13: make([]byte, size) escapes to heap
>./main.go:95:13:        from make([]byte, size) (too large for stack) at ./main.go:95:13

>报告显示，对于栈来说，数组太大了。这个信息误导了我们。并不是说底层的数组太大，而是编译器在编译时并不知道数组的大小。
>值只有在编译器编译时知道其大小才会将它分配到栈中。这是因为每个函数的栈帧大小是在编译时计算的。
>如果编译器不知道其大小，就只会在堆中分配。

>为了验证（我们的想法），我们将值硬编码为 5，然后再次运行压力测试。
`buf := make([]byte, 5)`

>这一次我们运行压力测试，分配消失了

`go test -run none -bench AlgorithmTwo -benchtime 3s -benchmem -memprofile mem.out`

`BenchmarkAlgorithmOne-8    3000000      1720 ns/op        0 B/op        0 allocs/op`

 `go build -gcflags "-m -m"`

`./stream.go:95: algOne &bytes.Buffer literal does not escape`

`./stream.go:95: algOne make([]byte, 5) does not escape`
