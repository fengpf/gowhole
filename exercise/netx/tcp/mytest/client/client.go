package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	if err != nil {
		panic(err)
	}

	// fmt.Fprintf(conn, "i am client"+"\n")
	// // conn.Write([]byte("i am client\n"))
	// var buf [255]byte
	// conn.Read(buf[:])
	// fmt.Println(222, string(buf[:]))

	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		conn.Write([]byte(text + "\n"))
		// fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}

}
