package tests

import (
	"fmt"
	"testing"
	"time"
)

//golang 的返回值是通过栈空间，不是通过寄存器，这点最重要。
//调用函数前，首先分配的是返回值空间，然后是入参地址，再是其他临时变量地址。
//return 操作  1、将返回值拷贝到栈空间第一块区域 2、判断defer函数是否修改栈空间的返回值 3、空的return（ret 跳转）
func Test_defer(t *testing.T) {
	fmt.Println(f(), f2(), f3())

	i := 100
	j := 200
	defer fmt.Printf("start i=%d, j=%d\n", i, j)
	defer func(v int) {
		defer fmt.Printf("end v=%d, i=%d, j=%d\n", v, i, j)
	}(j)
	i = 300
	j = 400
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// Golang中defer、return、返回值之间执行顺序的坑 by  henrylee2cn

func Test_defer2(t *testing.T) {
	fmt.Println("a return:", a()) // 打印结果为 a return: 0
	fmt.Println("b return:", b()) // 打印结果为 b return: 2
	c := c()
	fmt.Println("c return:", *c, c) // 打印结果为 c return: 2 0xc42008e378
	defer P(time.Now())
	time.Sleep(5e9)
	fmt.Println("1", time.Now())
}

//A. 匿名返回值的情况
func a() int {
	var i int
	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	return i
}

//B. 有名返回值的情况
func b() (i int) {
	defer func() {
		i++
		fmt.Println("b defer2:", i) // 打印结果为 b defer2: 2
	}()
	defer func() {
		i++
		fmt.Println("b defer1:", i) // 打印结果为 b defer1: 1
	}()
	return i // 或者直接 return 效果相同
}

// 1.多个defer的执行顺序为“后进先出”；
// 2.所有函数在执行RET返回指令之前，都会先检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回；
// 3.匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值；
// 4.return其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）；第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；
// ‍‍因此，‍‍return、defer、返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。

// C. 下面我们再来看第三个例子，验证上面的结论
func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i, &i) // 打印结果为 c defer2: 2 0xc42008e378
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i, &i) // 打印结果为 c defer1: 1 0xc42008e378
	}()
	return &i
}

// 虽然 c()*int 的返回值没有被提前声明，但是由于 c()*int 的返回值是指针变量，那么在return将变量 i 的地址赋给返回值后，defer再次修改了 i 在内存中的实际值，因此return调用RET退出函数时返回值虽然依旧是原来的指针地址，但是其指向的内存实际值已经被成功修改了。
// 即，我们假设的结论是正确的！

// D.defer声明时会先计算确定参数的值，defer推迟执行的仅是其函数体。
func P(t time.Time) {
	fmt.Println("2", t)
	fmt.Println("3", time.Now())
}

// 输出结果：
// 1 2018-07-11 11:59:40.157970811 +0800 CST m=+5.002160808
// 2 2018-07-11 11:59:35.156940164 +0800 CST m=+0.001130161
// 3 2018-07-11 11:59:40.158185175 +0800 CST m=+5.002375172

// E. defer的作用域

// 1. defer只对当前协程有效（main可以看作是主协程）；
// 2. 当panic发生时依然会执行当前（主）协程中已声明的defer，但如果所有defer都未调用recover()进行异常恢复，则会在执行完所有defer后引发整个进程崩溃；
// 3. 主动调用os.Exit(int)退出进程时，已声明的defer将不再被执行。
