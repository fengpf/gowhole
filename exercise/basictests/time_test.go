package basictests

import (
	"fmt"
	"testing"
	"time"
)

func Test_getCurTime(t *testing.T) {
	fmt.Println(time.Now().Unix())               //获取当前秒
	fmt.Println(time.Now().UnixNano())           //获取当前纳秒
	fmt.Println(time.Now().UnixNano() / 1e6)     //将纳秒转换为毫秒
	fmt.Println(time.Now().UnixNano() / 1e9)     //将纳秒转换为秒
	c := time.Unix(time.Now().UnixNano()/1e9, 0) //将毫秒转换为 time 类型
	fmt.Println(c.String())                      //输出当前英文时间戳格式
}

var (
	consumeLimit = int64(1)
	tokenChan    = make(chan int, consumeLimit) //速率缓冲大小控制
	consumeRate  = int64(1e9 / consumeLimit)
)

func Test_tick(t *testing.T) {
	a := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	go generateToken(time.Duration(consumeRate))
	for _, v := range a {
		token := <-tokenChan //消费速率控制
		fmt.Println(v, token)
	}
	time.Sleep(10 * time.Second)
}

func generateToken(duration time.Duration) {
	var (
		timer = time.NewTicker(duration)
		token = 0
	)
	for range timer.C {
		token++
		tokenChan <- token
	}
}
