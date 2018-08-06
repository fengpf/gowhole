package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	conn, err := l.Accept()
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(conn net.Conn) {
	// 	wg.Done()

	// conn.Write([]byte("i am server" + "\n"))
	for { //如果服务端一直等待读数据，不给客户端写数据，客户端读会阻塞等待
		var buf *bufio.Reader
		buf = bufio.NewReader(conn)
		msg, err := buf.ReadString('\n')
		if err == io.EOF {
			return
		}
		// fmt.Printf("buf.ReadString error(%v)\n", err)
		fmt.Print("Message from client: " + msg)
		newMsg := strings.ToUpper(msg)
		conn.Write([]byte(newMsg + "\n"))
	}

	// }(conn)
	// wg.Wait()
}
