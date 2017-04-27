package main

import (
	"net"
	"fmt"
	"gostudy/net/heartbeat"
)

func main() {
	heartbeat.CMap = make(map[string]*heartbeat.CS)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 6666, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	go heartbeat.PushGRT()
	heartbeat.Server(listen)
	select {}
}

