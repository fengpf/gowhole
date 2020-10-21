package test

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"

	"gowhole/middleware/gorpc/api"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

type Arith int

// Some of Arith's methods have value args, some have pointer args. That's deliberate.

func (t *Arith) Add(args Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

const RPCAddr = "127.0.0.1:9000"

var (
	newServer                 *rpc.Server
	serverAddr, newServerAddr string
	httpServerAddr            string
	once, newOnce, httpOnce   sync.Once
)

func StartServerByDefault() {
	rpc.Register(new(api.Arith))

	var l net.Listener
	l, serverAddr = listenTCP()
	log.Println("Test RPC server listening on", serverAddr)

	go rpc.Accept(l)
}

func StartServer() {
	server := rpc.NewServer()
	server.Register(new(api.Arith))

	var l net.Listener
	l, serverAddr = listenTCP()
	log.Println("Test RPC server listening on", serverAddr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return
		}
		go server.ServeConn(conn)
	}

}

func CallRPC(addr string) {
	client, err := Dial("tcp", addr)
	if err != nil {
		log.Fatal("dialing", err)
	}
	defer client.Close()

	// Synchronous calls
	args := &Args{7, 8}
	reply := new(Reply)
	err = client.Call(api.ServiceMethod, args, reply)
	if err != nil {
		fmt.Errorf("Add: expected no error but got string %q", err.Error())
	}
	if reply.C != args.A+args.B {
		fmt.Errorf("Add: expected %d got %d", reply.C, args.A+args.B)
	}

	log.Printf("CallRPC success reply(%v)", reply)
}

func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", RPCAddr) // any available address
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

// Dial connects to an RPC server at the specified network address.
func Dial(network, address string) (*rpc.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}
