package main

import (
	"fmt"
	"os"
	"sync/atomic"
	"syscall"
	"testing"
	"unsafe"
)

// 	mmap用于把文件映射到内存空间中，简单说mmap就是把一个文件的内容在内存里面做一个映像。映射成功后，用户对这段内存区域的修改可以直接反映到内核空间，
// 	同样，内核空间对这段区域的修改也直接反映用户空间。那么对于内核空间<---->用户空间两者之间需要大量数据传输等操作的话效率是非常高的。

// 2、基本函数
// 2.1 mmap func(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error)
//  fd ：文件描述符（由open函数返回）
//  offset：表示被映射对象（即文件）从那里开始对映，通常都是用0。 该值应该为大小为PAGE_SIZE的整数倍
//  start：要映射到的内存区域的起始地址，通常都是用NULL（NULL即为0）。NULL表示由内核来指定该内存地址
//  length：要映射的内存区域的大小
//  prot：期望的内存保护标志，不能与文件的打开模式冲突。是以下的某个值，可以通过or运算合理地组合在一起
//  PROT_EXEC //页内容可以被执行
//  PROT_READ //页内容可以被读取
//  PROT_WRITE //页可以被写入
//  PROT_NONE //页不可访问

// flags：指定映射对象的类型，映射选项和映射页是否可以共享。它的值可以是一个或者多个以下位的组合体
// MAP_FIXED ：使用指定的映射起始地址，如果由start和len参数指定的内存区重叠于现存的映射空间，重叠部分将会被丢弃。如果指定的起始地址不可用，操作将会失败。并且起始地址必须落在页的边界上。
// MAP_SHARED ：对映射区域的写入数据会复制回文件内, 而且允许其他映射该文件的进程共享。
// MAP_PRIVATE ：建立一个写入时拷贝的私有映射。内存区域的写入不会影响到原文件。这个标志和以上标志是互斥的，只能使用其中一个。
// MAP_DENYWRITE ：这个标志被忽略。
// MAP_EXECUTABLE ：同上
// MAP_NORESERVE ：不要为这个映射保留交换空间。当交换空间被保留，对映射区修改的可能会得到保证。当交换空间不被保留，同时内存不足，对映射区的修改会引起段违例信号。
// MAP_LOCKED ：锁定映射区的页面，从而防止页面被交换出内存。
// MAP_GROWSDOWN ：用于堆栈，告诉内核VM系统，映射区可以向下扩展。
// MAP_ANONYMOUS ：匿名映射，映射区不与任何文件关联。
// MAP_ANON ：MAP_ANONYMOUS的别称，不再被使用。
// MAP_FILE ：兼容标志，被忽略。
// MAP_32BIT ：将映射区放在进程地址空间的低2GB，MAP_FIXED指定时会被忽略。当前这个标志只在x86-64平台上得到支持。
// MAP_POPULATE ：为文件映射通过预读的方式准备好页表。随后对映射区的访问不会被页违例阻塞。
// MAP_NONBLOCK ：仅和MAP_POPULATE一起使用时才有意义。不执行预读，只为已存在于内存中的页面建立页表入口。

// 2.2 munmap func(addr uintptr, length uintptr) error
// start：要取消映射的内存区域的起始地址
// length：要取消映射的内存区域的大小。
// 返回说明
// 成功执行时munmap()返回0。失败时munmap返回-1.
// 对映射内存的内容的更改并不会立即更新到文件中，而是有一段时间的延迟，你可以调用msync()来显式同步一下, 这样你内存的更新就能立即保存到文件里
// start：要进行同步的映射的内存区域的起始地址。
// length：要同步的内存区域的大小
// flag:flags可以为以下三个值之一：
// MS_ASYNC : 请Kernel快将资料写入。
// MS_SYNC : 在msync结束返回前，将资料写入。
// MS_INVALIDATE : 让核心自行决定是否写入，仅在特殊状况下使用

// 3 用户空间和驱动程序的内存映射
// 3.1、基本过程
// 首先，驱动程序先分配好一段内存，接着用户进程通过库函数mmap()来告诉内核要将多大的内存映射到内核空间，
// 内核经过一系列函数调用后调用对应的驱动程序的file_operation中指定的mmap函数，
// 在该函数中调用remap_pfn_range()来建立映射关系。

// 3.2、映射的实现
// 首先在驱动程序分配一页大小的内存，然后用户进程通过mmap()将用户空间中大小也为一页的内存映射到内核空间这页内存上。
// 映射完成后，驱动程序往这段内存写10个字节数据，用户进程将这些数据显示出来。

func TestNonGoMemory(t *testing.T) {
	//Mmap(fd int, offset int64, length int, prot int, flags int)

	// fd, err := syscall.Socket(syscall.AF_LOCAL, syscall.SOCK_STREAM, 0)
	// if err != nil {
	// 	t.Fatalf("Socketpair: %v", err)
	// }
	// defer syscall.Close(fd)
	// writeFile := os.NewFile(uintptr(fd), "test writes")
	// defer writeFile.Close()

	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open shm file error(%v)", err)
		return
	}
	// f.WriteString("hello")
	fd := int(f.Fd())
	// fd, err := syscall.Open("./usdeamon.txt", syscall.O_RDWR, 0666)
	// if err != nil {
	// 	fmt.Printf("open shm file error(%v)", err)
	// 	return
	// }

	// info, err := f.Stat()
	// if err != nil {
	// 	fmt.Printf("f.Stat() error(%v)", err)
	// 	return
	// }
	// // Ensure the size is at least the minimum size.
	// var size = int(info.Size())
	// size, err = mmapSize(size)
	// if err != nil {
	// 	return
	// }

	err = syscall.Ftruncate(fd, int64(ps))
	if err != nil {
		fmt.Printf("syscall.Ftruncate error(%v)", err)
		return
	}
	data, err := syscall.Mmap(fd, 0, 3, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED) //test
	// data, err := syscall.Mmap(-1, 0, 3, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		t.Fatalf("failed to mmap memory: %v", err)
	}
	println(len(data), &data[0], &data[1], &data[2])

	//向内存区0存1，并且将存的值加1
	p := (*uint32)(unsafe.Pointer(&data[0]))
	fmt.Println(*p)
	atomic.AddUint32(p, 100)
	(*p)++
	fmt.Println(*p)
	if *p != 2 {
		t.Fatalf("data[0] = %v, expect 2", *p)
	}

	// p2 := (*uint32)(unsafe.Pointer(&data[1]))
	// atomic.AddUint32(p2, 3)

	fmt.Println(*p)
	syscall.Munmap(data)
	// fmt.Println(*p)
}

// 返回说明
// 成功执行时，mmap()返回被映射区的指针，munmap()返回0。失败时，mmap()返回MAP_FAILED[其值为(void *)-1]，munmap返回-1。errno被设为以下的某个值
// EACCES：访问出错
// EAGAIN：文件已被锁定，或者太多的内存已被锁定
// EBADF：fd不是有效的文件描述词
// EINVAL：一个或者多个参数无效
// ENFILE：已达到系统对打开文件的限制
// ENODEV：指定文件所在的文件系统不支持内存映射
// ENOMEM：内存不足，或者进程已超出最大内存映射数量
// EPERM：权能不足，操作不允许
// ETXTBSY：已写的方式打开文件，同时指定MAP_DENYWRITE标志
// SIGSEGV：试着向只读区写入
// SIGBUS：试着访问不属于进程的内存区

const (
	maxMmapStep = 1 << 30 // 1GB
	// maxMapSize represents the largest mmap size supported by Bolt.
	maxMapSize = 0x7FFFFFFF // 2GB
	// maxAllocSize is the size used when creating array pointers.
	maxAllocSize = 0xFFFFFFF
)

// mmapSize determines the appropriate size for the mmap given the current size
// of the database. The minimum size is 32KB and doubles until it reaches 1GB.
// Returns an error if the new mmap size is greater than the max allowed.
func mmapSize(size int) (int, error) {
	// Double the size from 32KB until 1GB.
	for i := uint(15); i <= 30; i++ {
		if size <= 1<<i {
			return 1 << i, nil
		}
	}

	// Verify the requested size is not above the maximum allowed.
	if size > maxMapSize {
		return 0, fmt.Errorf("mmap too large")
	}

	// If larger than 1GB then grow by 1GB at a time.
	sz := int64(size)
	if remainder := sz % int64(maxMmapStep); remainder > 0 {
		sz += int64(maxMmapStep) - remainder
	}

	// Ensure that the mmap size is a multiple of the page size.
	// This should always be true since we're incrementing in MBs.
	pageSize := int64(2)
	if (sz % pageSize) != 0 {
		sz = ((sz / pageSize) + 1) * pageSize
	}

	// If we've exceeded the max size then only grow up to the max size.
	if sz > maxMapSize {
		sz = maxMapSize
	}

	return int(sz), nil
}
