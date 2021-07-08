package main

import (
	"fmt"
	"net"
	"net/rpc"

	"gowhole/middleware/gorpc/api"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}

	req := api.Request{7, 8}
	reply := new(api.Reply)

	client := rpc.NewClient(conn)
	defer client.Close()

	err = client.Call(api.ServiceMethod, req, reply)
	if err != nil {
		fmt.Printf("client.Call error(%v)\n", err)
	}

	if reply.C == req.A+req.B {
		fmt.Printf("Add: expected %d got %d\n", reply.C, req.A+req.B)
	}

	req = api.Request{7, 0}
	reply = new(api.Reply)
	err = client.Call("Embed.Exported", req, reply)
	if err != nil {
		fmt.Printf("Add: expected no error but got string %v\n", err.Error())
	}
	if reply.C != req.A+req.B {
		fmt.Printf("Add: expected %d got %d\n", reply.C, req.A+req.B)
	}

	fmt.Println("rpc call success")
}
