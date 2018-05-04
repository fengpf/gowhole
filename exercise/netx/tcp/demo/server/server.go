package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	println("server start...")
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		if _, err := io.WriteString(conn, time.Now().Format("15:04:05\n")); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(60 * time.Second)
	}
}
