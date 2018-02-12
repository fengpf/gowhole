package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := context.TODO()
	url := "https://i0.hdslb.com/bfs/article/8fce42a26ce128140d2ee7dec599a46cd1bccbb6.jpg"
	capture(c, url)
}

func capture(c context.Context, url string) (loc string, size int, err error) {
	_, ct, err := download(c, url)
	if err != nil {
		return
	}
	println(url, ct)
	return loc, size, err
}

func download(c context.Context, url string) (bs []byte, ct string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("http.NewRequest %+v\n", err)
		return
	}
	// client
	client := &http.Client{
		Timeout: time.Duration(5),
	}
	// timeout
	// ctx, cancel := context.WithTimeout(c, 800*time.Millisecond)
	// req = req.WithContext(ctx)
	// defer cancel()
	//client do
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client.Do %+v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("resp.StatusCodeo %+v\n", resp.StatusCode)
		return
	}
	ct = resp.Header.Get("Content-Type")
	bs, _ = ioutil.ReadAll(resp.Body)
	ct = http.DetectContentType(bs)
	return
}
