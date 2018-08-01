package main

import (
	"fmt"
	"io"
	"os"

	"net/http"
)

var (
	filename = "a.txt"
	data     *string
)

func main() {
	var err error
	resp, err := http.DefaultClient.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// resp.Body 被读空,下面写文件没数据
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// f.Write(resp.Body)
	io.Copy(f, resp.Body) //写文件

	_, err = Contents(filename) //读取文件
	if err != nil {
		panic(err)
	}
	// fmt.Printf("out %s\n", out)
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
		// println(n)
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
