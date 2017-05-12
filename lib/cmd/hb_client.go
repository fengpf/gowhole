package main

import (
	"net"
	"fmt"
	"gostudy/lib/net/heartbeat"
)

func main() {
	heartbeat.Dch = make(chan bool)
	heartbeat.Rch = make(chan []byte)
	heartbeat.Wch = make(chan []byte)
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6666")
	println(addr.IP.String())
	conn, err := net.DialTCP("tcp", nil, addr)
	//	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	go heartbeat.ClientHandler(conn)
	select {
	case <-heartbeat.Dch:
		fmt.Println("关闭连接")
	}
}