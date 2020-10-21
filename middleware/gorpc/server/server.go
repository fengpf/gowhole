package main

import (
	"log"
	"net"
	"net/rpc"

	"gowhole/middleware/gorpc/api"
)

func main() {
	server := rpc.NewServer()

	var err error
	err = server.Register(new(api.Arith))

	//TODO match rpc.ServeConn
	//err = rpc.Register(new(api.Arith)) //bug: use rpc.Regiser will use default server,and can not find any service
	if err != nil {
		log.Fatalf("rpc.Register error(%v)\n", err)
	}

	var l net.Listener
	l, err = net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("net.Listen tcp error(%v)\n", err)
	}
	defer l.Close()

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Print("rpc.Serve: accept:", err.Error())
				return
			}
			go server.ServeConn(conn)

			//TODO match rpc.Register
			//go rpc.ServeConn(conn)
		}
	}()

	log.Printf("rpc start server address(%v)\n", l.Addr().String())
	select {}
}
