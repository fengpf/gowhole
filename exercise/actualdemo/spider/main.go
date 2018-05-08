package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

const (
	infoURL = "http://www.imooc.com"
)

type worker struct {
	job  chan int
	done func()
}

func hanle(http.ResponseWriter, *http.Request) {
	parse()
}

func main() {
	http.HandleFunc("/", hanle)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func parse() {
	req, err := http.NewRequest(http.MethodGet, infoURL, nil)
	// req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	client := http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", request)
			return nil
		},
	}
	// resp, err := http.DefaultClient.Do(req)//use default client
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

}

func doWork(i int, w worker) {
	for n := range w.job {
		fmt.Printf("Worker %d received %d\n", i, n)
		w.done()
	}
}

func genWorker(n int, wg *sync.WaitGroup) (wk worker) {
	wk = worker{
		job: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(n, wk)
	return
}

func dispatch() {
	var workers [10]worker
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		workers[i] = genWorker(i, &wg)
	}
	wg.Add(20)
	for i, w := range workers {
		w.job <- i //发完,stop发的要有人收
	}
	for i, w := range workers {
		w.job <- i + 1000 //stop 还没人就发，卡住
	}
	wg.Wait()
}
