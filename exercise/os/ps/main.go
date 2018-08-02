package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	c1 := exec.Command("ps", []string{"-ef"}...)
	stdout1, err := c1.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = c1.Start(); err != nil {
		log.Fatal(err)
	}
	c2 := exec.Command("grep", "brower")
	c2.Stdin = stdout1

	stdout2, err := c2.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err = c2.Start(); err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(stdout2)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
}
