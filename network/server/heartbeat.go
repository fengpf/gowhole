package main

import (
	"fmt"
	"net"
	"time"
)

//golang实现带有心跳检测的tcp长连接

var (
	// ReqREGISTER client register cid.
	ReqREGISTER byte = 1
	// ResREGISTER server response.
	ResREGISTER byte = 2
	// ReqHEARTBEAT send heartbeat req.
	ReqHEARTBEAT byte = 3
	// ResHEARTBEAT send heartbeat res.
	ResHEARTBEAT byte = 4
	// Req cs send data.
	Req byte = 5
	// Res cs send ack.
	Res byte = 6
)

// CS info.
type CS struct {
	Rch chan []byte
	Wch chan []byte
	Dch chan bool
	u   string
}

// NewCs new CS.
func NewCs(uid string) *CS {
	return &CS{
		Rch: make(chan []byte),
		Wch: make(chan []byte),
		u:   uid,
	}
}

// CMap CS map.
var CMap map[string]*CS

// PushGRT push.
func PushGRT() {
	for {
		time.Sleep(15 * time.Second)
		for k, v := range CMap {
			fmt.Println("push msg to user:" + k)
			v.Wch <- []byte{Req, '#', 'p', 'u', 's', 'h', '!'}
		}
	}
}

// Server listen.
func Server(listen *net.TCPListener) {
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("接受客户端连接异常:", err.Error())
			continue
		}
		fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
		go ServerHandler(conn)
	}
}

// ServerHandler server handle.
func ServerHandler(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 128)
	var uid string
	var C *CS
	for {
		conn.Read(data)
		fmt.Println("客户端发来数据:", string(data))
		if data[0] == ReqREGISTER { // register
			conn.Write([]byte{ResREGISTER, '#', 'o', 'k'})
			uid = string(data[2:])
			C = NewCs(uid)
			//fmt.Println("register client")
			//fmt.Println(uid)
			break
		} else {
			conn.Write([]byte{ResREGISTER, '#', 'e', 'r'})
		}
	}
	go ServerWHandler(conn, C)
	go ServerRHandler(conn, C)
	go ServerWork(C)
	select {
	case <-C.Dch:
		fmt.Println("close handler goroutine")
	}
}

// ServerWHandler 正常写数据.
func ServerWHandler(conn net.Conn, C *CS) {
	// 读取业务Work 写入Wch的数据
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case d := <-C.Wch:
			conn.Write(d)
		case <-ticker.C:
			if _, ok := CMap[C.u]; !ok {
				fmt.Println("conn die, close WHandler")
				return
			}
		}
	}
}

// ServerRHandler 正常读数据.
func ServerRHandler(conn net.Conn, C *CS) {
	// 心跳ack 业务数据 写入Wch
	for {
		data := make([]byte, 128)
		// setReadTimeout
		err := conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			fmt.Println(err)
		}
		if _, derr := conn.Read(data); derr == nil {
			// 可能是来自客户端的消息确认 数据消息
			fmt.Printf("服务端读取数据: %v\n", data)
			if data[0] == Res {
				fmt.Printf("服务端接收客户端 ack数据包: %v\n", string(data))
			} else if data[0] == Req {
				fmt.Printf("服务端读取正常请求数据: %v\n", data)
				conn.Write([]byte{Res, '#'})
				// C.Rch <- data
			}
			continue
		}
		conn.Write([]byte{ReqHEARTBEAT, '#'})
		fmt.Println("服务端发送心跳数据:")
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		if _, herr := conn.Read(data); herr == nil {
			fmt.Printf("服务端接收心跳数据: %v\n", string(data))
		} else {
			delete(CMap, C.u)
			fmt.Println("delete user!")
			return
		}
	}
}

// ServerWork Sleep.
func ServerWork(C *CS) {
	time.Sleep(5 * time.Second)
	C.Wch <- []byte{Req, '#', 'h', 'e', 'l', 'l', 'o'}

	time.Sleep(15 * time.Second)
	C.Wch <- []byte{Req, '#', 'h', 'e', 'l', 'l', 'o'}
	// 从读ch读信息
	/*	ticker := time.NewTicker(20 * time.Second)
		for {
			select {
			case d := <-C.Rch:
				C.Wch <- d
			case <-ticker.C:
				if _, ok := CMap[C.u]; !ok {
					return
				}
			}

		}
	*/ // 往写ch写信息
}

func main() {
	CMap = make(map[string]*CS)
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP("127.0.0.1"), 6666, ""})
	if err != nil {
		fmt.Println("监听端口失败:", err.Error())
		return
	}
	fmt.Println("已初始化连接，等待客户端连接...")
	go PushGRT()
	Server(listen)
	select {}
}
