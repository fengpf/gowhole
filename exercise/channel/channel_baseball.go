package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup
func init(){
	rand.Seed(time.Now().UnixNano())
}
func main(){
   ball:=make(chan int)
   wg.Add(2)
   go  play("a",ball)
   go  play("b",ball)
   ball<-1
   wg.Wait()
} 

func play(name string,ball chan int){
    defer wg.Done()
	for {
		counter,ok:=<-ball
		if !ok{
		  //被动关闭通道，表示赢球
		  fmt.Printf("%s won the ball\n",name)
		  return
		}
		//模拟丢球
		n:=rand.Intn(100)
		if n%13==0{
			fmt.Printf("%s miss %d\n",name, counter)
			close(ball)
			return
		}
		//模拟打球
		fmt.Printf("%s hit %d\n",name, counter)
		counter++
		ball<-counter
	}
}