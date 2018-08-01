package main

import (
	"fmt"
	"sync"
)

type message struct {
	next *message
	data int
	done bool
}

var (
	length     = 10
	chanSize   = 1024
	msgs       = make(chan *message)
	doneChan   = make(chan []*message, chanSize)
	mergeChan  = make([]chan *message, chanSize)
	head, last *message
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		i := 1
		defer wg.Done()
		for {
			msgs <- &message{data: i}
			i++
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					return
				}
				m := &message{data: msg.data}
				if head == nil {
					head = m
					last = m
				} else {
					last.next = m
					last = m
				}
				mergeChan[msg.data%length] <- m
			case done := <-doneChan:
				commits := make(map[int]string)
				for _, d := range done {
					d.done = true
				}
				for ; head != nil && head.done; head = head.next {
					fmt.Printf("commit  %d ================ is %v\n", head.data, head.done)
					commits[head.data] = "commit"
				}
				// for k, v := range commits {
				// 	fmt.Printf("%d is %s\n", k, v)
				// }
			}
		}
	}()

	for i := 0; i < length; i++ {
		c := make(chan *message, chanSize)
		mergeChan[i] = c
		wg.Add(1)
		go func(c chan *message) {
			defer wg.Done()
			var marked = make([]*message, 0, chanSize)
			for {
				select {
				case msg, ok := <-c:
					if !ok {
						return
					}
					fmt.Println("consume..................", msg)
					marked = append(marked, msg)
				}
				if len(marked) > 0 {
					doneChan <- marked
					fmt.Println("marked --------", len(marked))
					marked = make([]*message, 0, chanSize)
				}
			}
		}(c)
	}
	wg.Wait()
}
