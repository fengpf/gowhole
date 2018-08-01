package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

//sudo docker run -p 6379:6379   -d redis:latest redis-server --appendonly yes
var (
	redisClient *redis.Pool
	host        = "127.0.0.1:6379"
	pwd         = "123"
)

func init() {
	// 建立连接池
	redisClient = &redis.Pool{
		MaxIdle:     5,                 //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   20,                //最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: 180 * time.Second, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) { //建立连接
			c, err := redis.Dial("tcp", host)
			if err != nil {
				fmt.Printf("redis.Dial error(%v)\n", err)
				return nil, err
			}
			// if _, err := c.Do("AUTH", pwd); err != nil {
			// 	c.Close()
			// 	fmt.Printf("c.Do auth error(%v)\n", err)
			// 	return nil, err
			// }
			// c.Do("SELECT", 0)//选择db
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

type msg struct {
	MID       int64 `json:"mid"`
	From      int   `json:"from"`
	IsAuthor  int   `json:"is_author"`
	Timestamp int64 `json:"timestamp"`
}

func main() {
	// concurrentBySameMID()

	concurrentByDifferentMID()
	a := make(chan bool, 1)
	<-a
}

func concurrentBySameMID() {
	mid := int64(27515211)
	for i := 0; i < 15; i++ { //同一个mid的并发
		data := &msg{
			MID:       mid,
			From:      0,
			IsAuthor:  i % 3,
			Timestamp: time.Now().Unix(),
		}
		go func(mid int64, data interface{}) {
			pub(mid, data)
		}(mid, data)
	}
}

func concurrentByDifferentMID() {
	mids := []int64{111, 222, 333, 444, 555}
	for i, mid := range mids { //不同mid的并发
		data := &msg{
			MID:       mid,
			From:      0,
			IsAuthor:  i % 3,
			Timestamp: time.Now().Unix(),
		}
		go func(mid int64, data interface{}) {
			pub(mid, data)
		}(mid, data)
	}
}

func pub(mid int64, data interface{}) {
	var (
		err  error
		jd   []byte
		conn redis.Conn
	)
	conn = redisClient.Get() // 从池里获取连接
	if conn == nil {
		return
	}
	jd, err = json.Marshal(data)
	if err != nil {
		fmt.Printf("json.Marshal error(%v)\n", err)
		return
	}
	err = conn.Send("SET", mid, jd)
	if err != nil {
		fmt.Printf("conn.Send error(%v)\n", err)
		return
	}
	fmt.Println("set success", string(jd))
	// 用完后将连接放回连接池
	defer conn.Close()
}
