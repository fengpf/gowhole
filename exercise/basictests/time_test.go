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

func Test_tick(t *testing.T) {
	var timer = time.NewTicker(time.Duration(1) * time.Second)
	var token = 0
	go func() {
		for range timer.C {
			token++
			fmt.Println(token)
			if token == 20 {
				break
			}
		}
	}()
	time.Sleep(5 * time.Second)
}
