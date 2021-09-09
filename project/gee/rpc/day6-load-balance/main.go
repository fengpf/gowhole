package main

import (
	"context"
	"gowhole/project/gee/rpc"
	"gowhole/project/gee/rpc/discovery"
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)
	ch1 := make(chan string)
	ch2 := make(chan string)
	// start two servers
	go startServer(ch1)
	go startServer(ch2)

	addr1 := <-ch1
	addr2 := <-ch2

	time.Sleep(time.Second)
	call(addr1, addr2)
	broadcast(addr1, addr2)
}

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}

func startServer(addrCh chan string) {
	var foo Foo
	l, _ := net.Listen("tcp", ":0")
	server := rpc.NewServer()
	_ = server.Register(&foo)
	addrCh <- l.Addr().String()
	server.Accept(l)
}

//封装一个方法 foo，便于在 Call 或 Broadcast 之后统一打印成功或失败的日志。
func foo(xc *discovery.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, args, &reply)
	case "broadcast":
		err = xc.Broadcast(ctx, serviceMethod, args, &reply)
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err)
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply)
	}
}

//call 调用单个服务实例，broadcast 调用所有服务实例
func call(addr1, addr2 string) {
	d := discovery.NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := discovery.NewXClient(d, discovery.RandomSelect, nil)
	defer func() { _ = xc.Close() }()
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "call", "Foo.Sum", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait()
}

func broadcast(addr1, addr2 string) {
	d := discovery.NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
	xc := discovery.NewXClient(d, discovery.RandomSelect, nil)

	defer func() { _ = xc.Close() }()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			foo(xc, context.Background(), "broadcast", "Foo.Sum", &Args{Num1: i, Num2: i * i})
			// expect 2 - 5 timeout
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait()
}
