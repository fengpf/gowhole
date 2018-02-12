package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file := "test.txt"
	out, err := Contents(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("out %s\n", out)
}

// Contents get file content.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	var res []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		println(n)
		// println(string(buf[0:n]))
		res = append(res, buf[0:n]...)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(res), nil // f will be closed if we return here.
}
