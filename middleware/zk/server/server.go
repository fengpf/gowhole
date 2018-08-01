package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"gowhole/lib/zk/util"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	go starServer("127.0.0.1:2181")
	go starServer("127.0.0.1:2182")
	go starServer("127.0.0.1:2183")

	a := make(chan bool, 1)
	<-a
}

func starServer(addr string) {
	var (
		err      error
		conn     *zk.Conn
		tcpAddr  *net.TCPAddr
		listener *net.TCPListener
	)
	if tcpAddr, err = net.ResolveTCPAddr("tcp4", addr); err != nil {
		panic(err)
	}

	if listener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		panic(err)
	}

	if conn, err = util.GetConnect(); err != nil {
		fmt.Printf("connect zk error(%v)\n", err)
		return
	}
	defer conn.Close()

	if err = util.RegistServer(conn, addr); err != nil { //注册zk节点
		fmt.Printf("regist node error(%v)\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "listener.Accept error(%v)", err)
			continue
		}
		go handleCient(conn, addr)
	}
}

func handleCient(conn net.Conn, port string) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(port + ": " + daytime))
}
