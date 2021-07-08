package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"time"

	"gowhole/lib/zk/util"
)

func main() {
	for i := 0; i < 10; i++ {
		startClient()
		time.Sleep(1 * time.Second)
	}
}

func startClient() {
	var (
		err        error
		serverHost string
		tcpAddr    *net.TCPAddr
		conn       *net.TCPConn
	)

	if serverHost, err = getServerHost(); err != nil {
		panic(err)
	}

	if tcpAddr, err = net.ResolveTCPAddr("tcp4", serverHost); err != nil {
		panic(err)
	}

	if conn, err = net.DialTCP("tcp", nil, tcpAddr); err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("timestamp"))
	result, err := ioutil.ReadAll(conn)
	fmt.Println(string(result))
	return
}

func getServerHost() (host string, err error) {
	conn, err := util.GetConnect()
	defer conn.Close()
	if err != nil {
		fmt.Printf("util.GetConnect() error(%v)\n", err)
		return
	}

	serverList, err := util.GetServerList(conn)
	if err != nil {
		fmt.Printf("util.GetServerList error(%v)\n", err)
		return
	}

	count := len(serverList)
	if count == 0 {
		return "", errors.New("server list is empty")
	}
	fmt.Printf("serverList(%v)\n", serverList)

	//随机选中一个返回
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	host = serverList[r.Intn(3)]
	return
}
