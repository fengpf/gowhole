package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

var (
	err error
	ps  = syscall.Getpagesize()
)

func main() {
	// mmap()
	// regSingal()

	var (
		p     []byte
		_p0   unsafe.Pointer
		n     int
		_zero uintptr //Single-word zero for use when we need a valid pointer to 0 bytes.
	)
	if len(p) > 0 {
		_p0 = unsafe.Pointer(&p[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}

	fd, err := syscall.Open("./t.txt", syscall.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("syscall.Open error(%v)", err)
		return
	}

	r0, _, e1 := syscall.Syscall(syscall.SYS_READ, uintptr(fd), uintptr(_p0), uintptr(len(p)))
	n = int(r0)
	if e1 != 0 {
		println(e1.Error())
	}

	println(n, _zero, _p0, string(p[:]))

}

//mmap 操作共享内存
func mmap() {
	fd, err := syscall.Open("./test.txt", syscall.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("syscall.Open error(%v)", err)
		return
	}
	err = syscall.Ftruncate(fd, int64(ps))
	if err != nil {
		fmt.Printf("syscall.Ftruncate error(%v)", err)
		return
	}
	bs, err := syscall.Mmap(fd, 0, ps, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Printf("syscall.Mmap error(%v)", err)
		return
	}

	// a := (*string)(unsafe.Pointer(&bs[0]))
	// *a = "hello"
	// println(*a)

	in := "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest"
	inByte := []byte(in)
	copy(bs, inByte)

	fmt.Println(string(bs), ps)
	// os.Exit(-1)
}

func regSingal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		syscall.SIGBUS,  //总线错误
		syscall.SIGABRT, //调用abort
		syscall.SIGSEGV, //访问无效内存
		syscall.SIGILL,  //栈溢出
		//测试用
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGSTOP,
		syscall.SIGINT,
		syscall.SIGHUP,
	)

	// 忽略掉其他信号
	signal.Ignore(syscall.SIGHUP, //终端挂起
		syscall.SIGPIPE) //管道错误
	for {
		c := <-ch
		switch c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			fmt.Println("mmap save.....")
			mmap()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
