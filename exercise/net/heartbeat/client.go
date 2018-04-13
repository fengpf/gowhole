package heartbeat

// golang实现带有心跳检测的tcp长连接
// client

import (
	"fmt"
	"net"
)

// Dch detect chan.
var Dch chan bool

// Rch read chan.
var Rch chan []byte

// Wch write chan.
var Wch chan []byte

// ClientHandler client.
func ClientHandler(conn *net.TCPConn) {
	// 直到register ok
	data := make([]byte, 128)
	for {
		conn.Write([]byte{ReqREGISTER, '#', '2'})
		conn.Read(data)
		//fmt.Println(string(data))
		if data[0] == ResREGISTER {
			break
		}
	}
	//fmt.Println("i'm register")
	go ClinetRHandler(conn)
	go ClinetWHandler(conn)
	go ClinetWork()
}

// ClinetRHandler read handler.
func ClinetRHandler(conn *net.TCPConn) {
	for {
		// 心跳包,回复ack
		data := make([]byte, 128)
		i, _ := conn.Read(data)
		if i == 0 {
			Dch <- true
			return
		}
		if data[0] == ReqHEARTBEAT {
			fmt.Printf("客户端发送心跳数据包: %v\n", string(data))
			conn.Write([]byte{ResREGISTER, '#', 'h'})
		} else if data[0] == Req {
			fmt.Printf("客户端接收正常数据包: %v\n", string(data))
			fmt.Printf("%v\n", string(data[2:]))
			Rch <- data[2:]
			conn.Write([]byte{Res, '#'})
		}
	}
}

// ClinetWHandler write handler.
func ClinetWHandler(conn net.Conn) {
	for {
		select {
		case msg := <-Wch:
			fmt.Println((msg[0]))
			fmt.Println("send data after: " + string(msg[1:]))
			conn.Write(msg)
		}
	}

}

// ClinetWork recv data.
func ClinetWork() {
	for {
		select {
		case msg := <-Rch:
			fmt.Println("work recv " + string(msg))
			Wch <- []byte{Req, '#', 'x', 'x', 'x', 'x', 'x'}
		}
	}
}
